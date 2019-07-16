<?php

function test() {\PhpBench::startTimer(1);


    sleep(2);
\PhpBench::timeCode(1, __FILE__, 4, trim('sleep(2);'));
\PhpBench::startTimer(2);


    echo "done";
\PhpBench::timeCode(2, __FILE__, 6, trim('echo "done";'));

}\PhpBench::startTimer(3);


test();
\PhpBench::timeCode(3, __FILE__, 9, trim('test();'));
