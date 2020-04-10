package main

import "fmt"

func main() {

	for t := 1; t <= 8; t++ {
		for h := 1; h <= 8; h++ {
			if h != t {
				for b := 2; b <= 8; b++ {
					if b != t && b != h {
						for s := 1; s <= 7; s++ {
							if s != b && s != h && s != t {
								m := s + 1
								for r := 1; r <= 8; r++ {
									if r != m && r != s && r != t && r != h && r != b {
										for i := 1; i <= 8; i++ {
											if i != r && i != m && i != s && i != t && i != h && i != b {
												for c := 1; c <= 8; c++ {
													if c != i && c != r && c != m && c != s && c != t && c != h && c != b {
														the := 100*t + 10*h
														best := 1000*b + 10*s + t
														system := 100000*s + 90000 + 1000*s + 100*t + m
														metric := 100000*m + 1000*t + 100*r + 10*i + c
														if (the + best + system) == metric {
															fmt.Println("     THE")
															fmt.Println("    BEST")
															fmt.Println("+ SYSTEM")
															fmt.Println("--------")
															fmt.Println("= METRIC")
															fmt.Println()
															fmt.Printf("     %d\n", the)
															fmt.Printf("    %d\n", best)
															fmt.Printf("+ %d\n", system)
															fmt.Println("--------")
															fmt.Printf("= %d\n", metric)
															fmt.Println()
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
