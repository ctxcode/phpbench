<?php

if (!class_exists('PhpBench')) {
    class PhpBench {

        static $decreaseTime = null;
        static $setNr = null;
        static $timers = [];

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
}
