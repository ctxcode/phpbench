<?php

class PhpBench {

    static $timer = null;

    public static function startTimer() {
        static::$timer = microtime(true);
    }

    public static function timeCode($code) {
        $end = microtime(true);

        $timeSpent = $end-static::$timer;

        $ms = round($timeSpent * 1000);

        echo $ms . 'ms';

        static::startTimer();
    }

}

?>