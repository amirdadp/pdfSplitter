#!/usr/bin/env bash



begin=1
end=1
filename=${1}
counter=1
shift


for x in ${@} ; do

	end=${x}
	gs -sDEVICE=pdfwrite -dNOPAUSE -dBATCH -dSAFER -dFirstPage=${begin} -dLastPage=${end} -sOutputFile="part${counter}.pdf"  ${filename}

	counter=$(($counter + 1 )) 
	begin=$(($end + 1 )) 
done



