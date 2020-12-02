package com.mikeseagull.adventofcode2017;

import java.util.List;

public class Part2 extends Part1 {
    @Override
    public int main() {
        List<Integer> instructions = readFile("part2_input");
        int tracker = 0;
        int step_counter = 0;
        while (tracker < instructions.size()) {
            step_counter++;
            int original_tracker = tracker;
            int instruction = instructions.get(tracker);
            tracker += instruction;
            if (instruction >= 3) {
                instructions.set(original_tracker, instruction - 1);
            } else {
                instructions.set(original_tracker, instruction + 1);
            }
        }
        return step_counter;
    }
}
