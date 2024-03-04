#!/usr/bin/env bash
# leehom Chen clh021@gmail.com
cd "$( dirname "${BASH_SOURCE[0]}" )" || exit
sqlite3 test.db ".output db.dump.sql" ".dump"
