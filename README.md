# Aim -  企业客服IM中台

## 概述：
aim是一个即时通讯服务器，主要为客服平台服务。实现以下功能：
1. 支持http、tcp、websocket接入
2. 消息记录云端存储，离线消息同步
3. 支持微信客服接入
4. 支持水平拓展
5. 支持emoji表情、多媒体文件、微信资讯卡片
6. 客服主动接入、转接
7. 自动回复

## 环境要求：
1. [Go] >= 1.14
2. [Nsq]
3. [Redis]
4. [Mysql]

## 依赖：
1. [Grpc] - RPC服务（Protocol Buffers）
2. [Gorm] - ORM框架
3. [Redigo] — Redis工具
4. [Zap] - 日志工具
5. [Wiidz/Goutil] - go工具包