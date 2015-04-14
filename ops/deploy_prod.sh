#!/bin/bash

ansible-playbook deploy_prod.yml -i inventory/production.hosts --vault-password-file ~/.vault_pass.txt