CREATE TABLE `user` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `username` varchar(20) DEFAULT NULL,
    `phone` varchar(20) DEFAULT NULL,
    `email` varchar(20) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;


CREATE TABLE `t_employee` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `username` varchar(20) DEFAULT NULL,
    `name` varchar(20) DEFAULT NULL,
    `age` int(20) DEFAULT NULL,
    `create_time` datetime NOT NULL DEFAULT current_timestamp(),
    `edit_time` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;