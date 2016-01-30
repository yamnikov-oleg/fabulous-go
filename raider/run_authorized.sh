echo -n "GitHub username: "
read name

echo -n "GitHub password: "
read -s pass
echo

go build
export GITHUB_USERNAME=$name
export GITHUB_PASSWORD=$pass
./raider
