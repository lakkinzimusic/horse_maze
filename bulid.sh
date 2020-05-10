gitUsername=$(git config user.name)
gitPassword=$(git config user.password)

git add .
git commit -m "message"
git push "https://${gitUsername}:${gitPassword}@github.com/lakkinzimusic/horse_maze.git"