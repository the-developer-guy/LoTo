# LoTo
Lockout-Tagout for your servers and services

## Goal
A [LoTo system](https://en.wikipedia.org/wiki/Lockout%E2%80%93tagout) is a hardware mutex. Its goal is to keep workers safe by showing "hey, someone has put a big red sticker on this switch" or physically preventing switches to be turned on.

If you go to the gym (or visit a nice beach), a towel can claim a bench or a seat. Nothing enforces it really, everyone is expected to play by the rules. This service is the towel.

The goal is to make a super simple LoTo for internal use. Do you have VMs or other shared services, which can only be used one person at a time? You probably already asked "Who is using this, is this free?" before. Just check the LoTo and claim it for yourself!

## Configuration

Mount the customized config.yaml (preferably to /etc/loto/config.yaml) and set the LOTO_CONFIG_PATH to the config file.
