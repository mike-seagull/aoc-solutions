#!/usr/bin/env node
const fs = require('fs');
const readline = require('readline');

function part1() {
	var lineReader = readline.createInterface({
	  input: fs.createReadStream('part1_input')
	});

	part1_valid_passphrases = 0

	lineReader.on('line', function (passphrase) {
	  var words = passphrase.split(" ");
	  var is_valid_passphrases = true;
	  while (words.length > 0) {
	  	var word = words.pop();
	  	if (words.indexOf(word) > -1) {
	  		is_valid_passphrases = false;
	  		break;
	  	}
	  }
	  if (is_valid_passphrases) {
	  	part1_valid_passphrases++;
	  }
	});
	lineReader.on('close', function() {
		console.log("part1 = ", part1_valid_passphrases);
	});
}

function part2() {
	var lineReader = readline.createInterface({
	  input: fs.createReadStream('part2_input')
	});

	part2_valid_passphrases = 0

	lineReader.on('line', function (passphrase) {
	  var words = passphrase.split(" ");
	  for (var i=0; i<words.length;i++) {
	  	words[i] = words[i].split('').sort().join('');
	  }
	  var is_valid_passphrases = true;
	  while (words.length > 0) {
	  	var word = words.pop();
	  	if (words.indexOf(word) > -1) {
	  		is_valid_passphrases = false;
	  		break;
	  	}
	  }
	  if (is_valid_passphrases) {
	  	part2_valid_passphrases++;
	  }
	});
	lineReader.on('close', function() {
		console.log("part2 = ", part2_valid_passphrases);
	});
}
part1();
part2();
