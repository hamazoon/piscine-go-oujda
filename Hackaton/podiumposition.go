package piscine

func PodiumPosition(podium [][]string) [][]string {
	if len(podium) == 0 {
		return podium
	}
	for i := 0; i < len(podium)/2; i++ {
		podium[i], podium[len(podium)-i-1] = podium[len(podium)-1-i], podium[i]
	}
	return podium
}
