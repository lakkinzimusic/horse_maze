
GITHUBUSER=$(git config --global user.name)
GITHUBPASSWORD=$(git config --global user.password)

# echo ${GITHUBPASSWORD}
git add .
git commit -m "update"
git push https://${GITHUBUSER}:${GITHUBPASSWORD}@horse_maze.git -all