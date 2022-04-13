

git checkout trunk

git pull

git add .

git commit -m "fix trunk"

git push -u origin trunk

git checkout master

git merge trunk

git push -u origin master

git checkout trunk

pause

