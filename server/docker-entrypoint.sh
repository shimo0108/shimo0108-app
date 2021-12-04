#!/bin/bash
set -e
# SSMにコンテナを登録
amazon-ssm-agent -register -code "${SSM_AGENT_CODE}" -id "${SSM_AGENT_ID}" -region "ap-northeast-1"
# バックグラウンド実行
amazon-ssm-agent &
exec "$@"
