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
        $result = eval($code);
    } catch (ParseError | Throwable $e) {
        // Report error somehow
        // $result = include $path;

        echo '<style>body,html{padding:0;margin:0;}</style>';
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

        echo '<pre style="margin:0 0 15px 0; padding: 10px; background-color:#eee;">';
        echo trim($code);
        echo '</pre>';
        echo '</div>';

        exit;
    }
    return $result;
}

function phpBenchGetCode($path) {

    $code = file_get_contents($path);

    $newCode = "";
    $lastWord = "";
    $inString = false;
    $stringEndChar = "";
    $timerCount = 0;
    $prevChar = "";
    $lastToken = "";
    $tokens = [];
    $allowNewCode = true;
    $allowNewCodeAfterNextBracket = false;
    $allowNewCodeAfterNextSemiColon = false;
    $lastCodeLine = "";
    $inCommentSingleLine = false;
    $inCommentMultiLine = false;
    $lineNr = 0;
    $foundPhpTag = false;

    $codeLength = strlen($code);
    for ($i = 0; $i < $codeLength; $i++) {

        $c = $code[$i];

        if ($c == "\n") {
            $lineNr++;
        }

        $lastCodeLine .= $c;

        if ($inString) {
            if ($c == $stringEndChar && $prevChar != "\\") {
                $inString = false;
                $stringEndChar = "";
            }
            $prevChar = $c;
            continue;
        }

        if ($inCommentMultiLine) {
            if ($c == "/" && $prevChar == "*") {
                $inCommentMultiLine = false;
            }
            $prevChar = $c;
            continue;
        }

        if ($inCommentSingleLine) {
            if ($c == "\n") {
                $inCommentSingleLine = false;
            }
            $prevChar = $c;
            continue;
        }

        if (ctype_alnum($c)) {
            $lastWord .= $c;
        } else {

            if ($lastWord == "php" && !$foundPhpTag) {
                // $newCode .= $lastCodeLine;
                $foundPhpTag = true;
                $lastCodeLine = "";
            }

            if ($lastWord == "namespace") {
                $allowNewCode = false;
                $allowNewCodeAfterNextSemiColon = true;
            }

            if (in_array($lastWord, ['class', 'function', 'if', 'else', 'elseif', 'for', 'while'], true)) {
                $lastToken = $lastWord;
                $allowNewCode = false;
                if ($lastWord != "class") {
                    $allowNewCodeAfterNextBracket = true;
                }
            }

            if ($c == "*" && $prevChar == "/") {
                $inCommentMultiLine = true;
            }

            if ($c == "/" && $prevChar == "/") {
                $inCommentSingleLine = true;
            }

            if ($c == "{") {
                $tokens[] = $lastToken;
                if ($allowNewCodeAfterNextBracket) {
                    $allowNewCodeAfterNextBracket = false;
                    $allowNewCode = true;
                }
                $newCode .= $lastCodeLine;
                $lastCodeLine = "";
            }

            if ($c == "}") {
                $allowNewCode = true;
                array_pop($tokens);
                if (count($tokens) > 0) {
                    $token = $tokens[count($tokens) - 1];
                    if ($token != "class") {
                        $allowNewCode = false;
                    }
                }
                $newCode .= $lastCodeLine;
                $lastCodeLine = "";
            }

            if ($c == "'" || $c == "\"") {
                $inString = true;
                $stringEndChar = $c;
            }

            if ($c == ";") {
                if ($allowNewCode) {
                    $timerCount++;
                    $newCode .= "\\PhpBench::startTimer(" . ($timerCount) . ");\n";
                }
                $newCode .= $lastCodeLine . "\n";
                if ($allowNewCode) {
                    $escCodeLine = str_replace("'", "", $lastCodeLine);
                    $escCodeLine = preg_replace('/\r?\n/', ' ', $escCodeLine);
                    $escCodeLine = trim($escCodeLine);
                    $newCode .= "\\PhpBench::timeCode(" . $timerCount . ", __FILE__, " . $lineNr . ", '" . $escCodeLine . "');\n";
                }
                if ($allowNewCodeAfterNextSemiColon) {
                    $allowNewCodeAfterNextSemiColon = false;
                    $allowNewCode = true;
                }
                $lastCodeLine = "";
            }

            // Reset lastWord
            $lastWord = "";
        }

        $prevChar = $c;
    }

    $newCode .= $lastCodeLine;

    // Create and replace constants
    $count = ++PhpBench::$constCounter;

    define('PHPBENCH__FILE__' . $count, $path);
    define('PHPBENCH__DIR__' . $count, dirname($path));

    $newCode = preg_replace('/([^a-zA-Z0-9_])__FILE__([^a-zA-Z0-9_])/', '$1PHPBENCH__FILE__' . $count . '$2', $newCode);
    $newCode = preg_replace('/([^a-zA-Z0-9_])__DIR__([^a-zA-Z0-9_])/', '$1PHPBENCH__DIR__' . $count . '$2', $newCode);

    return $newCode;
}
