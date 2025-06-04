#!/bin/zsh

echo "Are you ready to be great?"

your_change="run and study"

echo ${your_change}

unset your_change

echo $0 $1 $2 $?

:<<EOF
12341234
12341234
EOF

number=1
if test ${number} -eq 1
then
  echo "真"
fi

for day in 1 2 3 4 5
do
  echo "get it"
done

case ${number} in
  1) echo "你选择了正确的道路"
  ;;
esac
