#!/bin/bash

pkill redis-controller

sleep 2s

./bin/redis-controller

