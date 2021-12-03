
export function run(input: ArrayBuffer): ArrayBuffer {
	let inStr = String.UTF8.decode(input)
	let n = parseInt(inStr)
	let f = factorial(n as i32)
  
	let out = "factorial of " + inStr + "is " + f.toString()

	return String.UTF8.encode(out)
}

function factorial(n: i32): i32 {
	if (n == 1) { return 1 }

	return n * factorial(n-1)
}