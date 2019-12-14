#! /bin/bash
# データベースが存在しない場合作成
echo "CREATE DATABASE IF NOT EXISTS \`$MYSQL_DATABASE\` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;" | "${mysql[@]}"
# $MYSQL_DATADASEの権限を$MYSQL_USERに与える
echo "GRANT ALL ON \`$MYSQL_DATABASE\`.* TO '"$MYSQL_USER"'@'%' ;" | "${mysql[@]}"
# 権限の反映
echo 'FLUSH PRIVILEGES ;' | "${mysql[@]}"
