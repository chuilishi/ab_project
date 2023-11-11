CREATE TABLE `sys_users` (
    `id` bigint unsigned AUTO_INCREMENT,
    `created_at` datetime NULL,
    `updated_at` datetime NULL,
    `deleted_at` datetime NULL,
    `studentid` varchar(13) DEFAULT '未知',
    `username` varchar(5) DEFAULT '未知',
    `sex` varchar(2) DEFAULT '未知',
    `grade` varchar(15) DEFAULT '未知',
    `profession` varchar(30) DEFAULT '未知',
    `phone` varchar(11) DEFAULT '未知',
    `wxid` varchar(30) DEFAULT '未知',
    `direction` varchar(6) DEFAULT '未知',
    `wxopenid` varchar(50) DEFAULT '未知',
    `status` varchar(10) DEFAULT '未知',
    `introduction` varchar(300) DEFAULT'还没有自我介绍',
    `reasons` varchar(300) DEFAULT '未知',
    `experience` varchar(300) DEFAULT '未知',
    `award` varchar(300) DEFAULT '未知',
    `remark` varchar(300) DEFAULT '未知',
    `ok` bigint DEFAULT 0,
    `isproblem` bigint DEFAULT 0,
    `problem` varchar(300) DEFAULT '无异常信息',
    `first` varchar(300) DEFAULT '还没有安排',
    `second` varchar(300) DEFAULT '还没有结论',
    `third` varchar(300) DEFAULT '还没有安排',
    `fourth` varchar(300) DEFAULT '还没有结论',
    PRIMARY KEY (`id`,`wxopenid`),INDEX `idx_sys_users_deleted_at` (`deleted_at`))
