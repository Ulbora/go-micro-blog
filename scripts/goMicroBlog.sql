-- MySQL Script generated by MySQL Workbench
-- Tue 28 Feb 2023 08:41:33 PM EST
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema go_micro_blog
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema go_micro_blog
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `go_micro_blog` ;
USE `go_micro_blog` ;

-- -----------------------------------------------------
-- Table `go_micro_blog`.`role`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `go_micro_blog`.`role` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL DEFAULT 'user',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `go_micro_blog`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `go_micro_blog`.`user` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(100) NOT NULL,
  `password` VARCHAR(100) NULL,
  `first_name` VARCHAR(100) NOT NULL,
  `last_name` VARCHAR(100) NOT NULL,
  `image` LONGBLOB NULL,
  `role_id` BIGINT NOT NULL,
  `active` TINYINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE,
  INDEX `fk_user_role1_idx` (`role_id` ASC) VISIBLE,
  CONSTRAINT `fk_user_role1`
    FOREIGN KEY (`role_id`)
    REFERENCES `go_micro_blog`.`role` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `go_micro_blog`.`blog`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `go_micro_blog`.`blog` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NOT NULL,
  `content` BLOB(65000) NOT NULL,
  `user_id` BIGINT NOT NULL,
  `active` TINYINT NOT NULL DEFAULT 0,
  `entered` DATETIME NOT NULL,
  `updated` DATETIME NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_blog_user_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_blog_user`
    FOREIGN KEY (`user_id`)
    REFERENCES `go_micro_blog`.`user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `go_micro_blog`.`likes`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `go_micro_blog`.`likes` (
  `blog_id` BIGINT NOT NULL,
  `user_id` BIGINT NOT NULL,
  PRIMARY KEY (`blog_id`, `user_id`),
  INDEX `fk_likes_user1_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_likes_blog1`
    FOREIGN KEY (`blog_id`)
    REFERENCES `go_micro_blog`.`blog` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_likes_user1`
    FOREIGN KEY (`user_id`)
    REFERENCES `go_micro_blog`.`user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `go_micro_blog`.`comment`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `go_micro_blog`.`comment` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `blog_id` BIGINT NOT NULL,
  `user_id` BIGINT NOT NULL,
  `text` VARCHAR(500) NOT NULL,
  `active` TINYINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  INDEX `fk_comment_blog1_idx` (`blog_id` ASC) VISIBLE,
  INDEX `fk_comment_user1_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_comment_blog1`
    FOREIGN KEY (`blog_id`)
    REFERENCES `go_micro_blog`.`blog` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_comment_user1`
    FOREIGN KEY (`user_id`)
    REFERENCES `go_micro_blog`.`user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `go_micro_blog`.`user_auth`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `go_micro_blog`.`user_auth` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `type` VARCHAR(100) NOT NULL DEFAULT 'linkedin',
  `user_id` BIGINT NOT NULL,
  `date_entered` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_user_auth_user1_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_user_auth_user1`
    FOREIGN KEY (`user_id`)
    REFERENCES `go_micro_blog`.`user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `go_micro_blog`.`config`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `go_micro_blog`.`config` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `allow_auto_post` TINYINT NOT NULL DEFAULT 0,
  `allow_auto_comment` TINYINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;