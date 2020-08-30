cat<<EOF > temp
ehllo
EOF
s=`cat temp`
echo  $s

echo $s|mail -s '生日快樂' fishis@163.com 
