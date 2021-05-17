#!/usr/bin/python3

import re
from sys import argv


def main():
    print("M3U Downloader")

    filename = "input.json"

    if(len(argv) > 1):
        filename = argv[1]  # Read from args
    else:
        filename = input("Enter the M3U file name")
        if(len(filename) == 0):
            filename = 'playlist.m3u8'

    F = open(filename, "r")
    J = F.read()


main()
