<?php

class PhpBench {

    static $decreaseTime = null;
    static $setNr = null;
    static $timers = [];

    static $timerCount = 0;
    static $constCount = 0;

    public static function startTimer($timer) {
        static::$timers['t' . $timer] = microtime(true);
    }

    public static function timeCode($timer, $filename, $lineNr, $code) {
        $end = microtime(true);

        $start = static::$timers['t' . $timer];
        $timeSpent = $end - $start;
        unset(static::$timers['t' . $timer]);

        $ms = round($timeSpent * 1000 * 1000);

        if (!static::$setNr) {
            static::$setNr = floor(microtime(true) * 1000) . '';
        }

        if (!static::$decreaseTime) {
            static::$decreaseTime = floor($start / 10000) * 10000;
        }
        $start -= static::$decreaseTime;
        $end -= static::$decreaseTime;

        // echo $ms . 'ms';

        $url = 'http://127.0.0.1:3001/data';
        $data = [
            'filename' => $filename,
            'setNr' => static::$setNr,
            'key' => $timer,
            'lineNr' => $lineNr,
            'code' => $code,
            'ms' => $ms,
            'start' => floor($start * 1000 * 100),
            'end' => round($end * 1000 * 100),
        ];
        static::postWithoutWait($url, $data);
        return;

        // Get cURL resource
        $curl = curl_init();
        // Set some options - we are passing in a useragent too here
        curl_setopt_array($curl, [
            CURLOPT_RETURNTRANSFER => 1,
            CURLOPT_URL => $url,
            CURLOPT_POST => 1,
            CURLOPT_POSTFIELDS => $data,
        ]);
        // Send the request & save response to $resp
        $resp = curl_exec($curl);
        // Close request to clear up some resources
        curl_close($curl);

    }

    private static function postWithoutWait($url, $params) {
        foreach ($params as $key => &$val) {
            if (is_array($val)) {
                $val = implode(',', $val);
            }

            $post_params[] = $key . '=' . urlencode($val);
        }
        $post_string = implode('&', $post_params);

        $parts = parse_url($url);

        $fp = fsockopen($parts['host'],
            isset($parts['port']) ? $parts['port'] : 80,
            $errno, $errstr, 30);

        $out = "POST " . $parts['path'] . " HTTP/1.1\r\n";
        $out .= "Host: " . $parts['host'] . "\r\n";
        $out .= "Content-Type: application/x-www-form-urlencoded\r\n";
        $out .= "Content-Length: " . strlen($post_string) . "\r\n";
        $out .= "Connection: Close\r\n\r\n";
        if (isset($post_string)) {
            $out .= $post_string;
        }

        fwrite($fp, $out);
        fclose($fp);
    }

}

function phpbench_error($e, $code) {

    echo '<style>body,html{padding:0;margin:0;} pre span { display: block; } pre.phpbench-code span:before { counter-increment: line; content: counter(line); display: inline-block; min-width: 30px; margin-right: 5px; }</style>';
    echo '<div style="padding:15px;">';
    echo '<pre style="padding:10px 15px;  margin:0 0 10px 0; background-color:black; color:red;">phpbench seems to have problems understanding certain syntax, create an issue on our github, sorry 😥</pre>';

    echo '<pre style="margin:0 0 15px 0; padding: 10px; background-color:#eee;">';
    echo 'Error: ' . $e->getMessage() . "\n";
    echo 'File: ' . $e->getFile() . "\n";
    echo 'Line: ' . $e->getLine() . "\n";
    echo "\n";
    echo $e->getTraceAsString();
    echo '</pre>';

    echo '<pre style="padding:10px 15px; margin:0 0 10px 0; background-color:black; color:white;">The generated code</pre>';

    echo '<pre style="margin:0 0 15px 0; padding: 10px; background-color:#eee; counter-reset: line;" class="phpbench-code">';
    $lines = explode("\n", trim($code));
    foreach ($lines as $line) {echo '<span>' . $line . '</span>';}
    // echo trim($code);
    echo '</pre>';
    echo '</div>';

    exit;
}

function phpbench_include($path) {

    $fn = basename($path);
    if ($fn == 'autoload.php') {
        return $path;
    }

    if (!file_exists($path)) {
        throw \Exception('Trying to include file that doesnt exists: ' . $path);
    }

    ini_set('display_errors', 1);
    ini_set('display_startup_errors', 1);
    error_reporting(E_ALL);

    $code = phpBenchGetCode($path);

    // echo $path . '<br>';
    // if ($fn == 'init.php') {
    //     echo '<pre>';
    //     echo str_replace('<?php', '< ?php', $code);
    //     echo '</pre>';
    //     echo 'z';
    //     exit;
    // }

    $newPath = substr($path, 0, -4) . '.phpbench.php';
    file_put_contents($newPath, $code);

    return $newPath;

    // try {
    // $result = eval($code);
    // } catch (ParseError | Throwable $e) {
    //     phpbench_error($e);
    // }
    return $result;
}

function phpBenchGetCode($path) {

    $code = file_get_contents($path);

    $newCode = "";
    $lastCodeLine = "";
    $addCode = false;
    $lineNr = 1;
    $lastTokenWasNewLine = false;
    $firstPhpTag = true;

    $flushLastCodeLine = function () use (&$newCode, &$lastCodeLine) {
        $newCode .= $lastCodeLine;
        $lastCodeLine = "";
    };

    $tokens = token_get_all($code, TOKEN_PARSE);
    $count = count($tokens);

    $skipUntilScope = function (&$i) use (&$lastCodeLine, &$tokens, &$count, &$lineNr) {
        $bracks = 0;
        while ($i < $count) {
            $token = $tokens[$i];

            if ($token == "\n") {
                $lineNr++;
            }

            if ($token == '(') {
                $bracks++;
            }
            if ($token == ')') {
                $bracks--;
            }
            if ($bracks === 0 && ($token == '{' || $token == ";")) {
                $i--;
                break;
            }

            $lastCodeLine .= is_array($token) ? $token[1] : $token;

            $i++;
        }
    };

    $skipUntilSemi = function (&$i) use (&$lastCodeLine, &$tokens, &$count, &$lineNr) {
        $bracks = 0;
        while ($i < $count) {
            $token = $tokens[$i];

            if ($token == "\n") {
                $lineNr++;
            }

            if ($token == '(' || $token == '{') {
                $bracks++;
            }
            if ($token == ')' || $token == '}') {
                $bracks--;
            }
            if ($bracks === 0 && ($token == ";")) {
                $i--;
                break;
            }

            $lastCodeLine .= is_array($token) ? $token[1] : $token;

            $i++;
        }
    };

    for ($i = 0; $i < $count; $i++) {

        $token = $tokens[$i];

        if ($i > 1) {
            $lastTokenWasNewLine = false;
            $prevToken = $tokens[$i - 1];
            if (is_array($prevToken)) {
                $prevTokenName = token_name($prevToken[0]);
                if ($prevTokenName == 'T_WHITESPACE') {
                    if (strpos($prevToken[1], "\n") !== false) {
                        $prev2Token = $tokens[$i - 2];
                        if (in_array($prev2Token, [';', '{'], true)) {
                            $lastTokenWasNewLine = true;
                        }
                    }
                }
            }
        }

        if (is_array($token)) {
            $name = token_name($token[0]);

            if ($name == 'T_OPEN_TAG' && $firstPhpTag) {
                // $lastCodeLine .= '<?php';
                $firstPhpTag = false;
                // continue;
            }

            if (in_array($name, ['T_CLASS', 'T_FOR', 'T_FOREACH', 'T_WHILE', 'T_IF', 'T_ELSEIF', 'T_ELSE', 'T_FUNCTION', 'T_THROW', 'T_CATCH'], true)) {
                $skipUntilScope($i);
                continue;
            }

            if (in_array($name, ['T_VARIABLE', 'T_STRING', 'T_ECHO'], true)) {
                if (!$addCode && $lastTokenWasNewLine) {
                    $addCode = true;
                    $flushLastCodeLine();
                    $skipUntilSemi($i);
                    continue;
                }
            }

            $lastCodeLine .= $token[1];
            continue;
        }

        $lastCodeLine .= $token;

        if ($token == "\n") {
            $lineNr++;
        }

        if ($token == ';') {
            if ($addCode) {

                \PhpBench::$timerCount++;
                $newCode .= "\\PhpBench::startTimer(" . (\PhpBench::$timerCount) . ");\n";
                $newCode .= $lastCodeLine . "\n";

                $escCodeLine = str_replace("'", "", $lastCodeLine);
                $escCodeLine = preg_replace('/\r?\n/', ' ', $escCodeLine);
                $escCodeLine = trim($escCodeLine);
                $newCode .= "\\PhpBench::timeCode(" . \PhpBench::$timerCount . ", __FILE__, " . $lineNr . ", '" . $escCodeLine . "');\n";

                $lastCodeLine = "";
                $addCode = false;
                continue;
            }
        }

    }

    $flushLastCodeLine();

    $newCode .= $lastCodeLine;

    // Create and replace constants
    $count = ++\PhpBench::$constCount;

    define('PHPBENCH__FILE__' . $count, $path);
    define('PHPBENCH__DIR__' . $count, dirname($path));

    $newCode = preg_replace('/([^a-zA-Z0-9_])__FILE__([^a-zA-Z0-9_])/', '$1PHPBENCH__FILE__' . $count . '$2', $newCode);
    $newCode = preg_replace('/([^a-zA-Z0-9_])__DIR__([^a-zA-Z0-9_])/', '$1PHPBENCH__DIR__' . $count . '$2', $newCode);

    $newCode = preg_replace('/(\n *(?:return )?)(include|include_once|require|require_once)([^a-zA-Z0-9_][^;\n]+);/', '$1$2(phpbench_include($3));', $newCode);

    // echo '<pre>';
    // echo str_replace('<?php', '< ?php', $newCode);
    // echo '</pre>';
    // exit;

    return $newCode;
}
