
# PhpBench (WIP)

`phpbench` is a tool to see how many milliseconds each line of code takes within a certain file.

## Install

```
go get -u github.com/ctxkiwi/phpbench/phpbench
```

## Usage (example)

```
> cd /var/www/mywebsite
> phpbench app/controllers/SearchController.php
*phpbench opens webserver on http://localhost:3001/*
```
Now run your code. In this example, it's searching something on their website.
Now surf to: `http://localhost:3001/`
And you should see the results. You can keep testing while the server is running.

It works the same for simple cli scripts.
```
> phpbench ~/my-scripts/backup-script.php
*phpbench opens webserver on http://localhost:3001/*

Open a new terminal and run it.
> php ~/my-scripts/backup-script.php

And go check the results on http://localhost:3001/
```

## Disadvantages
Your original code will be moved to {myphpfile}.original.php. And the script you targeted will be replaced with new code. You can't make any changes in this file. You have to edit myphpfile.original.php if you want to update your code while the phpbench is running. Although the webserver will detect these changes, update itself and it won't throw your changes away afterwards.

If you're done testing. You need to press CTRL+C. It will detect this and move {myphpfile}.original.php back to {myphpfile}.php.

## Video example

[https://www.youtube.com/watch?v=-SqGel3BALU](https://www.youtube.com/watch?v=-SqGel3BALU)

