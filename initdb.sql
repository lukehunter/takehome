CREATE  TABLE IF NOT EXISTS `host` (
  `host_id` INT(11) NOT NULL AUTO_INCREMENT ,
  `ip` VARCHAR(256) NOT NULL ,
  PRIMARY KEY (`host_id`) ,
  UNIQUE INDEX `descriptor_UNIQUE` (`ip` ASC) )
ENGINE = InnoDB
AUTO_INCREMENT = 13
DEFAULT CHARACTER SET = utf8;

CREATE  TABLE IF NOT EXISTS `scan` (
  `scan_id` INT(11) NOT NULL AUTO_INCREMENT ,
  `host_id` INT(11) NOT NULL ,
  `date_performed` DATETIME NOT NULL ,
  `start_port` INT(11) NOT NULL ,
  `end_port` INT(11) NOT NULL ,
  `open_ports` BLOB NOT NULL ,
  PRIMARY KEY (`scan_id`) ,
  INDEX `date_performed_IDX` (`date_performed` DESC) ,
  INDEX `fk_scan_host_host_id_idx` (`host_id` ASC) ,
  CONSTRAINT `fk_scan_host_host_id`
    FOREIGN KEY (`host_id` )
    REFERENCES `host` (`host_id` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
AUTO_INCREMENT = 15
DEFAULT CHARACTER SET = utf8;