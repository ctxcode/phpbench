<?php

class PhpBench {

    static $setNr = null;
    static $timers = [];

    public static function startTimer($timer) {
        static::$timers['t' . $timer] = microtime(true);
    }

    public static function timeCode($timer, $filename, $lineNr, $code) {
        $end = microtime(true);

        $timeSpent = $end-static::$timers['t' . $timer];
        unset(static::$timers['t' . $timer]);

        $ms = round($timeSpent * 1000 * 1000);

        if (!static::$setNr) {
            static::$setNr = floor(microtime(true) * 1000) . '';
        }

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
            ],
        ]);
        // Send the request & save response to $resp
        $resp = curl_exec($curl);
        // Close request to clear up some resources
        curl_close($curl);

    }

}
