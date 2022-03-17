#!/bin/bash

pkill -9 redis-controller

sleep 2s


nohup ./bin/redis-controller > ./logs/redis-controller.log &