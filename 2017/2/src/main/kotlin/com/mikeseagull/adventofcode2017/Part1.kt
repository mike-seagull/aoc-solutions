package com.mikeseagull.adventofcode2017

open class Part1 {

    fun readFile(fileName: String): ArrayList<String> {
        var inputStream = this.javaClass::class.java.getResource(fileName).openStream()
        val lineList = mutableListOf<String>()
        inputStream.bufferedReader().useLines {
            lines -> lines.forEach {
            lineList.add(it)
        }
        }
        var lines = ArrayList<String>()
        lineList.forEach {
            lines.add(it)
        }
        return lines
    }
    fun parseLine(line: String): ArrayList<Int> {
        var numbers = ArrayList<Int>()
        var beginIndex = 0
        var index = line.indexOf("\t")
        while (index >= 0) {
            numbers.add(Integer.parseInt(line.substring(beginIndex, index)))
            //println(index)
            beginIndex = index + 1
            if (line.indexOf("\t", index + 1) < 0) {
                break
            } else {
                index = line.indexOf("\t", index + 1)
            }
        }
        numbers.add(Integer.parseInt(line.substring(index + 1)))
        return numbers
    }
    fun differenceOfMaxAndMin(numbers: ArrayList<Int>): Int {
        numbers.sortDescending()
        var max: Int = numbers.get(0)
        var min: Int = numbers.last()
        return max - min
    }
    open fun main(): Int {
        var lines = readFile("/part1_input")
        var checksum = 0
        for (line in lines) {
            var numbers = parseLine(line)
            var difference = differenceOfMaxAndMin(numbers)
            checksum = checksum + difference
        }
        return checksum
    }
}