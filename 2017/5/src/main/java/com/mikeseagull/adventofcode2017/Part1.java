package com.mikeseagull.adventofcode2017;

import java.io.File;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Part1 {
    protected List<Integer> readFile(String fileName) {
        List<Integer> instructions = new ArrayList<Integer>();

        File inputFile = new File(this.getClass().getClassLoader().getResource(fileName).getFile());
        try (Scanner scanner = new Scanner(inputFile)) {
            while (scanner.hasNextLine()) {
                String line = scanner.nextLine();
                instructions.add(Integer.parseInt(line.trim()));
            }
            scanner.close();
        } catch (IOException e) {}
        return instructions;
    }
    public int main() {
        List<Integer> instructions = readFile("part1_input");
        int tracker = 0;
        int step_counter = 0;
        while (tracker < instructions.size()) {
            step_counter++;
            int original_tracker = tracker;
            int instruction = instructions.get(tracker);
            tracker += instruction;
            instructions.set(original_tracker, instruction + 1);
        }
        return step_counter;
    }
}
