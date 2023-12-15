#!/bin/bash

day=$(LANG=ja_JP.UTF-8 date +%a)

if grep -q "${day}" ./release.txt; then
  echo "今日はリリース日です"
else
  echo "今日はリリース日ではありません"
fi
