#!/bin/bash

pkill -9 redis-controller

sleep 2s

cd ./bin

nohup ./redis-controller > ../logs/redis-controller.log &

cd -