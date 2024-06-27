<?php 

$assoc_array = [
    "a"=>".-",
    "b"=>"-...", 
    "c"=>"-.-.", 
    "d"=>"-..", 
    "e"=>".", 
    "f"=>"..-.", 
    "g"=>"--.", 
    "h"=>"....", 
    "i"=>"..", 
    "j"=>".---", 
    "k"=>"-.-", 
    "l"=>".-..", 
    "m"=>"--", 
    "n"=>"-.", 
    "o"=>"---", 
    "p"=>".--.", 
    "q"=>"--.-", 
    "r"=>".-.", 
    "s"=>"...", 
    "t"=>"-", 
    "u"=>"..-", 
    "v"=>"...-", 
    "w"=>".--", 
    "x"=>"-..-", 
    "y"=>"-.--", 
    "z"=>"--..", 
    "0"=>"-----",
    "1"=>".----", 
    "2"=>"..---", 
    "3"=>"...--", 
    "4"=>"....-", 
    "5"=>".....", 
    "6"=>"-....", 
    "7"=>"--...", 
    "8"=>"---..", 
    "9"=>"----.",
    "."=>".-.-.-",
    ","=>"--..--",
    "?"=>"..--..",
    "/"=>"-..-.",
    " "=>" "];

    $morseCode = ".--. ..-  -.-  ..-";

    // Split words
    $reverted = '';
    $revertedCol = [];
    $explodedWords = explode("  ", $morseCode);
    foreach ( $explodedWords as $xword ) {
        $explodedLetters = explode(" ", $xword);
        foreach ( $explodedLetters as $xletter ) {
            $key = findKey($assoc_array, $xletter);
            $reverted .= (empty($key) ? " " : $key);
        }
    }

    echo "Reverted: $reverted\n";


    function findKey( $chunk, $morseCode ) {
        foreach ( $chunk as $key => $mcode ) {
            if ( $mcode == $morseCode ) {
                return $key;
            }
        }
        return false;
    }
    