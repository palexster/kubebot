#!/usr/bin/env bash

bash <( curl -s -H "Authorization: token ${GITHUB_TOKEN}" -H 'Accept: application/vnd.github.v3.raw' -L https://api.github.com/repos/gesundheitscloud/hc-deployment/contents/remote-scripts/deploy.sh )