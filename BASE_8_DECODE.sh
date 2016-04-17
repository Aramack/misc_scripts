# Input: (stdin) A white space seperated string that is base 8 encoded.
# Output: (stdout) The ASCII representation of the string
while read line
do
  echo "${line}" | sed -e 's/ /\n/g' | xargs -I % echo -e '\0%' | xargs echo -n | sed -e 's/ //g'
done < "${1:-/dev/stdin}"
