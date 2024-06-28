package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type NodeAM struct {
	text       string
	tag        string
	attributes map[string]string
	isChild    bool
	childs     []NodeAM
}

func parseArrow(arrowStr string) ([]NodeAM, error) {
	re := regexp.MustCompile(`(<-[^>]*->|-[^-]*-)|([^<\-]+)`)
	nodes := re.FindAllString(arrowStr, -1)
	nodesAM := make([]NodeAM, 0, len(nodes))
	for _, node := range nodes {
		nodesAM = append(nodesAM, NodeAM{
			text: node,
		})
	}

	for i, nodeAM := range nodesAM {
		if isNodeAMTag(nodeAM.text) {
			insides := strings.Split(nodeAM.text, " ")

			if nodeAM.text[:2] == "-<" {
				tag := insides[0][2:]

				if tag[len(tag)-1] == '-' {
					tag = tag[:len(tag)-1]
				}

				if i == 0 {
					return []NodeAM{}, errors.New("wrong usage of tag")
				}
				for j := i - 1; j >= 0; j-- {
					if isNodeAMTag(nodesAM[j].text) {
						if nodesAM[j].isChild {
							nodesAM[i].childs = append(nodesAM[i].childs, nodesAM[j])
						} else {
							fmt.Println(nodesAM[j])
							break
						}
					}
					nodesAM[j].tag = tag
				}

			} else {
				tag := insides[0][1:]
				if tag[len(tag)-2:] == ">-" {
					tag = tag[:len(tag)-2]
				}

				if i == len(nodesAM) {
					return []NodeAM{}, errors.New("wrong usage of tag")
				}

				for j := i + 1; j < len(nodesAM); j++ {
					if isNodeAMTag(nodesAM[j].text) {
						if nodesAM[j].isChild {
							nodesAM[i].childs = append(nodesAM[i].childs, nodesAM[j])
						} else {
							break
						}
					}
					nodesAM[j].tag = tag
				}
			}
		}
	}

	return nodesAM, nil
}

func isNodeAMTag(str string) bool {
	re := regexp.MustCompile(`^(-<\w+-|-[\w\s]+>-)$`)
	return re.MatchString(str)
}
