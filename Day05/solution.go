package main

import (
    "fmt"
    "strings"
    "utils"
)

type Rule struct {
    first  int
    second int
}

func extractRules(rules []string) []Rule {
    extractedRules := make([]Rule, len(rules))
    for i, rule := range rules {
        r := strings.Split(rule, "|")
        extractedRules[i] = Rule{utils.Stoi(r[0]), utils.Stoi(r[1])}
    }
    return extractedRules
}

func extractQueue(queue []string) [][]int {
    extractedQueue := make([][]int, len(queue))
    for i, line := range queue {
        numbers := strings.Split(line, ",")
        extractedNumbers := make([]int, len(numbers))
        for j, num := range numbers {
            extractedNumbers[j] = utils.Stoi(num)
        }
        extractedQueue[i] = extractedNumbers
    }
    return extractedQueue
}

func getIndex(value int, queue []int) int {
    for i, v := range queue {
        if v == value {
            return i
        }
    }
    return -1
}

func checkPrintingOrder(rules []Rule, printQueue []int) bool {
    for _, rule := range rules {
        secondIndex := getIndex(rule.second, printQueue)
        firstIndex := getIndex(rule.first, printQueue)
        if secondIndex != -1 && firstIndex > secondIndex {
            return false
        }
    }
    return true
}

func reorderQueue(rules []Rule, queue []int) []int {
    reordered := make([]int, len(queue))
    copy(reordered, queue)

    for {
        changed := false
        for _, rule := range rules {
            firstIndex := getIndex(rule.first, reordered)
            secondIndex := getIndex(rule.second, reordered)
            if firstIndex != -1 && secondIndex != -1 && firstIndex > secondIndex {
                reordered[firstIndex], reordered[secondIndex] = reordered[secondIndex], reordered[firstIndex]
                changed = true
            }
        }
        if !changed {
            break
        }
    }
    return reordered
}

func main() {
    dataRules, queue := utils.ReadFileWithEmptyLines("input.txt")
    rules := extractRules(dataRules)
    printQueue := extractQueue(queue)
    sumCorrect, sumIncorrect := 0, 0

    for _, pq := range printQueue {
        if checkPrintingOrder(rules, pq) {
            sumCorrect += pq[len(pq)/2]
        } else {
            reordered := reorderQueue(rules, pq)
            sumIncorrect += reordered[len(reordered)/2]
        }
    }

    fmt.Printf("Part 1: %d\nPart 2: %d\n", sumCorrect, sumIncorrect)
}