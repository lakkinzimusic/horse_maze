gitUsername=$(git config user.name)
gitPassword=$(git config user.password)
echo "${gitPassword}"

# echo "https://github.com/"${gitUsername}":"${gitPassword}"/horse_maze.git --all"
git add .
git commit -m "message"
git push "https://{lakkinzimusic}:{ubanoz55555}@github.com/lakkinzimusic/horse_maze.git --all"