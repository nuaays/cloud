
go mod graph | sed -Ee 's/@[^[:blank:]]+//g' | sort | uniq >unver.txt
cat unver.txt | awk '{print "\""$1"\" -> \""$2"\""};' >>graph. dot
echo "}" >>graph.dot
sed -i '' 's+\("github.com/[^/]*/\)\([^"]*"\)+\1\\n\2+g' graph. dot

