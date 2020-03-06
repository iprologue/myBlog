

CREATE TABLE `blog_tag` (
    `id` int(10) unsigned NOT NULL  AUTO_INCREMENT,
    `name` varchar(100) DEFAULT '' COMMENT '标签名称',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `created_by` varchar(100) default '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) default '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0',
    `state` tinyint(3) unsigned DEFAULT '1'  COMMENT '状态 0禁用、 1启用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT  CHARSET = utf8 COMMENT ='文章标签页';



CREATE TABLE `blog_article` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
    `title` varchar(100) DEFAULT '' COMMENT '文章标题',
    `desc` varchar(255) DEFAULT  '' COMMENT '简述',
    `content` text COMMENT '内容',
    `created_on` int(11) DEFAULT '0' COMMENT '创建日期',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0禁用1启用',
    PRIMARY KEY (`id`)
)ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT = '文章管理';