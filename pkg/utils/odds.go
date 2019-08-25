package utils

var Odds = []float64{
	1.01, 1.02, 1.03, 1.04, 1.05, 1.06, 1.07, 1.08, 1.09, 1.10,
	1.11, 1.12, 1.13, 1.14, 1.15, 1.16, 1.17, 1.18, 1.19, 1.20,
	1.21, 1.22, 1.23, 1.24, 1.25, 1.26, 1.27, 1.28, 1.29, 1.30,
	1.31, 1.32, 1.33, 1.34, 1.35, 1.36, 1.37, 1.38, 1.39, 1.40,
	1.41, 1.42, 1.43, 1.44, 1.45, 1.46, 1.47, 1.48, 1.49, 1.50,
	1.51, 1.52, 1.53, 1.54, 1.55, 1.56, 1.57, 1.58, 1.59, 1.60,
	1.61, 1.62, 1.63, 1.64, 1.65, 1.66, 1.67, 1.68, 1.69, 1.70,
	1.71, 1.72, 1.73, 1.74, 1.75, 1.76, 1.77, 1.78, 1.79, 1.80,
	1.81, 1.82, 1.83, 1.84, 1.85, 1.86, 1.87, 1.88, 1.89, 1.90,
	1.91, 1.92, 1.93, 1.94, 1.95, 1.96, 1.97, 1.98, 1.99, 2.00,
	2.02, 2.04, 2.06, 2.08, 2.10, 2.12, 2.14, 2.16, 2.18, 2.20,
	2.22, 2.24, 2.26, 2.28, 2.30, 2.32, 2.34, 2.36, 2.38, 2.40,
	2.42, 2.44, 2.46, 2.48, 2.50, 2.52, 2.54, 2.56, 2.58, 2.60,
	2.62, 2.64, 2.66, 2.68, 2.70, 2.72, 2.74, 2.76, 2.78, 2.80,
	2.82, 2.84, 2.86, 2.88, 2.90, 2.92, 2.94, 2.96, 2.98, 3.00,
	3.05, 3.10, 3.15, 3.20, 3.25, 3.30, 3.35, 3.40, 3.45, 3.50,
	3.55, 3.60, 3.65, 3.70, 3.75, 3.80, 3.85, 3.90, 3.95, 4.00,
	4.1, 4.2, 4.3, 4.4, 4.5, 4.6, 4.7, 4.8, 4.9, 5.0, 5.1, 5.2,
	5.3, 5.4, 5.5, 5.6, 5.7, 5.8, 5.9, 6.0, 6.2, 6.4, 6.6, 6.8,
	7.0, 7.2, 7.4, 7.6, 7.8, 8.0, 8.2, 8.4, 8.6, 8.8, 9.0, 9.2,
	9.4, 9.6, 9.8, 10.0, 10.5, 11, 11.5, 12, 12.5, 13, 13.5, 14,
	14.5, 15, 15.5, 16, 16.5, 17, 17.5, 18, 18.5, 19, 19.5, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100,
	110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 210, 220,
	230, 240, 250, 260, 270, 280, 290, 300, 310, 320, 330, 340,
	350, 360, 370, 380, 390, 400, 410, 420, 430, 440, 450, 460,
	470, 480, 490, 500, 510, 520, 530, 540, 550, 560, 570, 580,
	590, 600, 610, 620, 630, 640, 650, 660, 670, 680, 690, 700,
	710, 720, 730, 740, 750, 760, 770, 780, 790, 800, 810, 820,
	830, 840, 850, 860, 870, 880, 890, 900, 910, 920, 930, 940,
	950, 960, 970, 980, 990, 1000}

var OddsRange = []map[string]float64{
	map[string]float64{"begin": 1.01, "end": 2, "var": 0.01, "ticks": 100},
	map[string]float64{"begin": 2.02, "end": 3, "var": 0.02, "ticks": 50},
	map[string]float64{"begin": 3.05, "end": 4, "var": 0.05, "ticks": 20},
	map[string]float64{"begin": 4.1, "end": 6, "var": 0.1, "ticks": 20},
	map[string]float64{"begin": 6.2, "end": 10, "var": 0.2, "ticks": 20},
	map[string]float64{"begin": 10.5, "end": 20, "var": 0.5, "ticks": 20},
	map[string]float64{"begin": 21, "end": 30, "var": 1, "ticks": 10},
	map[string]float64{"begin": 32, "end": 50, "var": 2, "ticks": 10},
	map[string]float64{"begin": 55, "end": 100, "var": 5, "ticks": 10},
	map[string]float64{"begin": 110, "end": 1000, "var": 10, "ticks": 90},
}

func OddFloor(odd float64) float64 {
	// Boundaries
	if odd >= 1000 {
		return 1000
	} else if odd <= 1.01 {
		return 1.01
	}

	// Change this for a bisect version.
	for i := len(Odds) - 1; i >= 0; i-- {
		if Odds[i] <= odd {
			return Odds[i]
		}
	}

	return 0
}

func OddCeil(odd float64) float64 {
	// Boundaries
	if odd >= 1000 {
		return 1000
	} else if odd <= 1.01 {
		return 1.01
	}

	// Change this for a bisect version.
	for _, value := range Odds {
		if value >= odd {
			return value
		}
	}

	// Never hits this
	return 0
}

// def get_odd_shift(odd, shift):
//     """Calculate the odd after shift.

//     Args:
//         odd (float): odd to be shifted.
//         shift (int): shift of the odd (positive increases odd).

//     Returns: odd shifted.

//     """
//     try:
//         index = Odds.ODDS.index(odd)
//     except ValueError:
//         raise ValueError("Odd does not exist.")

//     index += shift

//     if index < 0 or index > Odds.ODDSLEN:
//         raise ValueError("Odds Ladder: Index [{}/{}] out of range."
//                          .format(index, Odds.ODDSLEN))

//     return Odds.ODDS[index]

// def odd_round(odd):
//     # TODO:   - Optimize the search algorithm. Bissection
//     dist = None
//     index = None

//     for key, value in enumerate(Odds.ODDS):
//         temp = abs(odd - value)

//         if dist is None or temp < dist:
//             dist = temp
//             index = key

//     return Odds.ODDS[index]
