#!/bin/sh
#
#       /etc/rc.d/init.d/todos
#
#       todos daemon
#
# chkconfig:   2345 95 05
# description: a todos script

### BEGIN INIT INFO
# Provides:       todos
# Required-Start:
# Required-Stop:
# Should-Start:
# Should-Stop:
# Default-Start: 2 3 4 5
# Default-Stop:  0 1 6
# Short-Description: todos
# Description: todos
### END INIT INFO

cd "$(dirname "$0")"

test -f .env && . $(pwd -P)/.env

_start() {
    test $(ulimit -n) -lt 100000 && ulimit -n 100000
    (env ENV=${ENV:-development} ./todos) <&- >todos.error.log 2>&1 &
    local pid=$!
    echo -n "Starting todos(${pid}): "
    sleep 1
    if (ps ax 2>/dev/null || ps) | grep "${pid} " >/dev/null 2>&1; then
        echo "OK"
    else
        echo "Failed"
    fi
}

_stop() {
    local pid="$(pidof todos)"
    echo -n "Stopping todos(${pid}): "
    if kill -HUP ${pid}; then
        echo "OK"
    else
        echo "Failed"
    fi
}

_restart() {
    if ! ./todos -validate ${ENV:-development}.toml >/dev/null 2>&1; then
        echo "Cannot restart todos, please correct todos toml file"
        echo "Run './todos -validate' for details"
        exit 1
    fi
    _stop
    sleep 1
    _start
}

_usage() {
    echo "Usage: [sudo] $(basename "$0") {start|stop|restart}" >&2
    exit 1
}

_${1:-usage}
