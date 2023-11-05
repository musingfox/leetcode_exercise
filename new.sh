#!/bin/zsh

original_text=$1
text_without_dot=${original_text//./}
final_text=${text_without_dot// /_}

echo $final_text
mkdir "$final_text"
