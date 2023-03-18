tanggal=$(date)
system_name=$(uname -a)
username=$(whoami)
test_ping=$(ping -c 3 google.com)
read name
read instagram
read linkedin
mkdir "$name $tanggal"
mkdir {"$name $tanggal"/about_me,"$name $tanggal"/my_friends,"$name $tanggal"/my_system_info}
mkdir {"$name $tanggal"/about_me/personal,"$name $tanggal"/about_me/professional}
touch "$name $tanggal"/about_me/personal/instagram.txt
touch "$name $tanggal"/about_me/professional/linkedin.txt
touch "$name $tanggal"/my_friends/list_of_my_friends.txt
touch "$name $tanggal"/my_system_info/about_this_laptop.txt
touch "$name $tanggal"/my_system_info/internet_connection.txt
echo "https://www.instagram.com/$instagram/" > "$name $tanggal"/about_me/personal/instagram.txt
echo "https://www.linkedin.com/in/$linkedin/" > "$name $tanggal"/about_me/professional/linkedin.txt
curl -o "$name $tanggal"/my_friends/list_of_my_friends.txt https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt
echo "My username: $username" >>"$name $tanggal"/my_system_info/about_this_laptop.txt
echo "With host: $system_name" >> "$name $tanggal"/my_system_info/about_this_laptop.txt
echo "$test_ping" > "$name $tanggal"/my_system_info/internet_connection.txt
