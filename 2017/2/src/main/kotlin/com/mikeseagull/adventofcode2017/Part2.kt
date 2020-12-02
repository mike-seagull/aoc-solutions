package com.mikeseagull.adventofcode2017

import java.util.*
import kotlin.collections.ArrayList

class Part2: Part1() {
    fun evenlyDivisible(checkNumber: Int, numbers: ArrayList<Int>): Int {
        for (number in numbers) {
            if (number > checkNumber && number % checkNumber == 0) {
                //return Pair(number, checkNumber)
                return number / checkNumber
            } else if (checkNumber % number == 0) {
                return checkNumber / number
            }
        }
        var number = numbers.get(0)
        numbers.removeAt(0)
        return evenlyDivisible(number, numbers)
    }
    override fun main(): Int {
        var lines = readFile("/part2_input")

        var checksum = 0
        for (line in lines) {
            var numbers = parseLine(line)
            var number = numbers.get(0)
            numbers.removeAt(0)
            checksum = checksum + evenlyDivisible(number, numbers)
        }
        return checksum
    }
}