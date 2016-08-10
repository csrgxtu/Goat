#!/usr/local/env python
# coding=utf-8
#
# Author: Archer
# File: wordCount.py
# Desc: 统计日志里面的词频
# Date: 10/Aut/2016
import urllib

RAW_WORDS = []
with open('/Users/archer/Downloads/keywords.txt') as F:
    for line in F:
        RAW_WORDS.append(urllib.unquote(line.strip('\n')))

with open('/Users/archer/Downloads/results.txt', 'w') as W:
    for w in RAW_WORDS:
        W.write(w)
        W.write('\n')
