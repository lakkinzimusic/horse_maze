
GITHUBUSER=$(git config --global user.name)
GITHUBPASSWORD=$(git config --global user.password)

git add .
git commit -m "update"
git push 