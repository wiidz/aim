# 「身份」
1. 聊天双方分客户（user）和职员（staff）

# 「内容」
1. 回复和接收支持文字、图片、视频、卡片讯息
2.

# 「记录」
1. 聊天记录尽可能的长时间备份（一个单位时间期间后，打包固化）
2. 接待历史也需要被记录
3. 


# 「群组」
1. 用户在进入聊天系统后，默认会创建一个只有自己一个人人的群组，接入客服后变为2人，可以无感知的过度到多人（这个设计主要是为了能共享聊天记录）
2.

# 「接入和结束」
1. 用户发起新聊天时，首先会将他的专属客服拉入群组，若无则将本次会话推入接待池中
1. 最后一次客户发言超过10分钟后无回应，推入应答池
1. 接入的客服会自动加入到用户的群组中
1. 可以手动结束本次会话
1. （待定）打分？
1. 结束聊天时，群组中会剔除除了用户以外的所有人
1. 

# 「转接和退出」
1. staff可以将本次会话转接给另一个staff，转接成功后新的staff加入群组，旧的staff退出
2. 当本次会话中的staff人数>=2人时，其中一人可以手动退出，不再接收消息

# 「数据」
1. 需要集成数据统计功能