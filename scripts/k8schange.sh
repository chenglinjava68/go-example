#!/bin/bash
#����䶯�󣬰汾û�䣬k8s������ȡ������Ҫ�ı�һ�������
sed -i "/- name: RAND_NUM/{ n;s/\(value: \).*/\1num$RANDOM/ }" $1.yaml
