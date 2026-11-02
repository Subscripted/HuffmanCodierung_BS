package main

import (
	"HuffmanCodierung_BS/types"
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var root *types.Node
	var bits string

	for {
		fmt.Println()
		fmt.Println("----- Modus Wählen -----")
		fmt.Println("1) Neuen Text codieren")
		fmt.Println("2) Letzten Text decodieren")
		fmt.Println("0) Beenden")
		fmt.Print("Auswahl: ")

		auswahl := leseZeile(reader)

		switch auswahl {
		case "1":
			fmt.Print("Text eingeben: ")
			text := leseZeile(reader)
			if text == "" {
				fmt.Println("Leerer Text kann nicht codiert werden.")
				continue
			}

			root = baueBaum(text)
			codes := make(map[rune]string)
			erzeugeCodes(root, "", codes)

			fmt.Println("Codes:")
			for r, c := range codes {
				fmt.Printf("  %q: %s\n", r, c)
			}

			bits = codiere(text, codes)
			fmt.Println("Codiert:", bits)
			fmt.Println(len(bits), "Bits statt", len(text)*8)

		case "2":
			if root == nil {
				fmt.Println("Es wurde noch kein Text codiert.")
				continue
			}
			fmt.Println("Decodiert:", decodiere(bits, root))

		case "0":
			fmt.Println("Tschüss!")
			return

		default:
			fmt.Println("Ungültige Auswahl.")
		}
	}
}

func leseZeile(reader *bufio.Reader) string {
	zeile, _ := reader.ReadString('\n')
	return strings.TrimRight(zeile, "\r\n")
}

func baueBaum(eingabetext string) *types.Node {

	if eingabetext == "" {
		return nil
	}

	runeMap := make(map[rune]int)
	for _, r := range eingabetext {
		runeMap[r]++
	}

	pq := &types.PriorityQueue{}
	for r, n := range runeMap {
		heap.Push(pq, &types.Node{Char: r, Frequency: n})
	}

	for pq.Len() > 1 {
		node1 := heap.Pop(pq).(*types.Node)
		node2 := heap.Pop(pq).(*types.Node)
		newnode := &types.Node{Left: node1, Right: node2, Frequency: node2.Frequency + node1.Frequency}
		heap.Push(pq, newnode)
	}

	root := heap.Pop(pq).(*types.Node)
	return root
}

func erzeugeCodes(node *types.Node, pfad string, codes map[rune]string) {
	//Node ein blatt? in codes für node.char aktuellen pdad ein,
	//Sonst rekursiv node.left, pdad + "0" und für node.right pfad + "1"

	if isLeaf(node) {
		if pfad == "" {
			pfad = "0"
		}
		codes[node.Char] = pfad
		return
	} else {
		erzeugeCodes(node.Left, pfad+"0", codes)
		erzeugeCodes(node.Right, pfad+"1", codes)
	}
}

func codiere(text string, codes map[rune]string) string {
	stringBuilder := strings.Builder{}
	for _, r := range text {
		stringBuilder.WriteString(codes[r])
	}
	return stringBuilder.String()
}

func decodiere(bits string, root *types.Node) string {
	stringBuilder := strings.Builder{}

	if isLeaf(root) {
		for range bits {
			stringBuilder.WriteRune(root.Char)
		}
		return stringBuilder.String()
	}

	node := root
	for _, r := range bits {
		if r == '0' {
			node = node.Left
		} else {
			node = node.Right
		}
		if isLeaf(node) {
			stringBuilder.WriteRune(node.Char)
			node = root
		}
	}
	return stringBuilder.String()
}

func isLeaf(node *types.Node) bool {
	return node.Left == nil && node.Right == nil
}
