param($exe, $inputDir, $outputDir)

# E.g. 
# > .\Make-AllDifferent.ps1 fivel.exe $env:DATA $env:DATA\FiveLetters\Go

$dictionaryFile = "$($inputDir)\british-english.txt"

$fiveLetterWordsFile = "$($outputDir)\five-letter-words.txt"

Start-Process $exe -NoNewWindow -Wait -ArgumentList 'find5LetterWords' -RedirectStandardInput $dictionaryFile -RedirectStandardOutput $fiveLetterWordsFile
