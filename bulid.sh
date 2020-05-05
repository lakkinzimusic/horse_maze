
GITHUBUSER=$(git config --global user.name)
GITHUBPASSWORD=$(git config --global user.password)

git remote set-url origin https://github.com/${GITHUBUSER}/horse_maze.git 
# echo ${GITHUBPASSWORD}
git add .
git commit -m "update"
# git push https://${GITHUBUSER}:${GITHUBPASSWORD}@github.com/horse_maze.git 
git push -u origin master