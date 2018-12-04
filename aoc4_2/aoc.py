#!/usr/bin/python3
from datetime import datetime

data = open('input').read().split('\n')


def line_to_slot(line):
    time = datetime.strptime(line[1:17], '%Y-%m-%d %H:%M')
    action = line[19:]
    return (time, action)


slots = [line_to_slot(line) for line in data]
slots = sorted(slots, key=lambda x: x[0])


def zero_hash():
    ret = {}
    for x in range(0, 60):
        ret[x] = 0
    return ret


guards = {}
guard = None
asleep_minute = None
asleep = False
for slot in slots:
    if '#' in slot[1]:
        asleep = False
        num_start = slot[1].index('#') + 1
        num_end = slot[1].index(' ', num_start)
        num = int(slot[1][num_start:num_end])
        if not num in guards:
            guards[num] = zero_hash()
        guard = num
    elif 'falls' in slot[1]:
        asleep_minute = slot[0].minute
        asleep = True
    elif 'wakes' in slot[1]:
        for minute in range(asleep_minute, slot[0].minute):
            guards[guard][minute] += 1

most_slept = -1
sleepy_guard = None
for guard, sleep in guards.items():
    slept = max(sleep.values())
    if slept > most_slept:
        print(guard, slept)
        most_slept = slept
        sleepy_guard = (guard, sleep)

slept_minute = -1
for minute, count in sleepy_guard[1].items():
    if count == most_slept:
        slept_minute = minute
        break

print(slept_minute, most_slept, sleepy_guard)
print(sleepy_guard[0] * slept_minute)
