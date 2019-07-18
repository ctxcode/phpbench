<?php

class PhpBench {

    static $decreaseTime = null;
    static $setNr = null;
    static $timers = [];

    static $constCounter = 0;

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

        // Get cURL resource
        $curl = curl_init();
        // Set some options - we are passing in a useragent too here
        curl_setopt_array($curl, [
            CURLOPT_RETURNTRANSFER => 1,
            CURLOPT_URL => 'http://127.0.0.1:3001/data',
            CURLOPT_POST => 1,
            CURLOPT_POSTFIELDS => [
                'filename' => $filename,
                'setNr' => static::$setNr,
                'key' => $timer,
                'lineNr' => $lineNr,
                'code' => $code,
                'ms' => $ms,
                'start' => floor($start * 1000 * 100),
                'end' => round($end * 1000 * 100),
            ],
        ]);
        // Send the request & save response to $resp
        $resp = curl_exec($curl);
        // Close request to clear up some resources
        curl_close($curl);

    }

}

function phpbench_include($path) {
    if (!file_exists($path)) {
        throw \Exception('Trying to include file that doesnt exists: ' . $path);
    }
    $code = phpBenchGetCode($path);

    try {
        $result = eval('?>' . $code);
    } catch (ParseError | Throwable $e) {
        // Report error somehow
        // $result = include $path;

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
    return $result;
}

function phpBenchGetCode($path) {

    $code = file_get_contents($path);

    $newCode = "";
    $lastCodeLine = "";
    $addCode = false;
    $lineNr = 1;
    $timerCount = 0;
    $lastTokenWasNewLine = false;

    $flushLastCodeLine = function () use (&$newCode, &$lastCodeLine) {
        $newCode .= $lastCodeLine;
        $lastCodeLine = "";
    };

    $tokens = token_get_all($code, TOKEN_PARSE);
    $count = count($tokens);

    $skipUntilScope = function (&$i) use (&$lastCodeLine, &$tokens, &$count) {
        $semicols = 0;
        $stop = false;
        while ($i < $count) {
            $token = $tokens[$i];

            if ($token == '(') {
                $semicols++;
            }
            if ($token == ')') {
                $semicols--;
            }
            if ($semicols === 0 && ($token == '{' || $token == ";")) {
                $stop = true;
            }

            $lastCodeLine .= is_array($token) ? $token[1] : $token;

            $i++;

            if ($stop) {
                $i--;
                break;
            }
        }
    };

    for ($i = 0; $i < $count; $i++) {

        $token = $tokens[$i];

        if ($i > 0) {
            $lastTokenWasNewLine = false;
            $prevToken = $tokens[$i];
            if (is_array($prevToken)) {
                $prevTokenName = token_name($prevToken[0]);
                if ($prevTokenName == 'T_WHITESPACE') {
                    if (strpos($prevToken[1], "\n") !== false) {
                        $lastTokenWasNewLine = true;
                    }
                }
            }
        }

        if (is_array($token)) {
            $name = token_name($token[0]);

            if ($name == 'T_OPEN_TAG') {
                $lastCodeLine .= '<?php';
                continue;
            }

            if (in_array($name, ['T_CLASS', 'T_FOR', 'T_FOREACH', 'T_WHILE', 'T_IF', 'T_ELSEIF', 'T_ELSE', 'T_FUNCTION', 'T_THROW', 'T_CATCH', 'T_CONST', 'T_PRIVATE', 'T_PUBLIC'], true)) {
                $skipUntilScope($i);
                continue;
            }

            if (in_array($name, ['T_VARIABLE', 'T_STRING'], true)) {
                if (!$addCode && $lastTokenWasNewLine) {
                    $addCode = true;
                    $flushLastCodeLine();
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

                $timerCount++;
                $newCode .= "\\PhpBench::startTimer(" . ($timerCount) . ");\n";
                $newCode .= $lastCodeLine . "\n";

                $escCodeLine = str_replace("'", "", $lastCodeLine);
                $escCodeLine = preg_replace('/\r?\n/', ' ', $escCodeLine);
                $escCodeLine = trim($escCodeLine);
                $newCode .= "\\PhpBench::timeCode(" . $timerCount . ", __FILE__, " . $lineNr . ", '" . $escCodeLine . "');\n";

                $lastCodeLine = "";
                $addCode = false;
                continue;
            }
        }

    }

    $flushLastCodeLine();

    $newCode .= $lastCodeLine;

    // Create and replace constants
    $count = ++PhpBench::$constCounter;

    define('PHPBENCH__FILE__' . $count, $path);
    define('PHPBENCH__DIR__' . $count, dirname($path));

    $newCode = preg_replace('/([^a-zA-Z0-9_])__FILE__([^a-zA-Z0-9_])/', '$1PHPBENCH__FILE__' . $count . '$2', $newCode);
    $newCode = preg_replace('/([^a-zA-Z0-9_])__DIR__([^a-zA-Z0-9_])/', '$1PHPBENCH__DIR__' . $count . '$2', $newCode);

    $newCode = preg_replace('/([^a-zA-Z0-9_])(include|include_once|require|require_once)([^a-zA-Z0-9_][^;\n]+);/', '$1phpbench_include($3);', $newCode);

    // echo '<pre>';
    // echo str_replace('<?php', '< ?php', $newCode);
    // echo '</pre>';
    // exit;

    return $newCode;
}
