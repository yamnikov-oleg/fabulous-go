read -p "GitHub username: " GITHUB_USERNAME
read -p "GitHub password: " -s GITHUB_PASSWORD
echo
export GITHUB_USERNAME
export GITHUB_PASSWORD
./raider
