-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jan 05, 2023 at 11:34 PM
-- Server version: 8.0.31-0ubuntu0.20.04.2
-- PHP Version: 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `imdbdbdemo`
--

-- --------------------------------------------------------

--
-- Table structure for table `movies`
--

CREATE TABLE `movies` (
  `id` bigint UNSIGNED NOT NULL,
  `moviecode` char(24) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `name` varchar(55) NOT NULL,
  `release_date` timestamp NOT NULL,
  `status` tinyint DEFAULT '0' COMMENT '1-Not Active, 2-Active, 3-Deactivated',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `person`
--

CREATE TABLE `person` (
  `id` int UNSIGNED NOT NULL,
  `name` varchar(65) NOT NULL,
  `email` varchar(255) NOT NULL,
  `mobile` varchar(255) NOT NULL,
  `age` tinyint DEFAULT NULL,
  `is_active` tinyint UNSIGNED DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `person`
--

INSERT INTO `person` (`id`, `name`, `email`, `mobile`, `age`, `is_active`, `created_at`, `updated_at`) VALUES
(63, 'user_x_w', 'n_s_user@gmail.com', '', 23, NULL, '2023-01-05 17:01:39', '2023-01-05 17:01:39'),
(64, 'user_c_p', 'l_t_user@gmail.com', '', 23, NULL, '2023-01-05 17:01:40', '2023-01-05 17:01:40'),
(65, 'user_q_m', 's_k_user@gmail.com', '', 23, NULL, '2023-01-05 17:01:41', '2023-01-05 17:01:41'),
(66, 'user_x_w', 'd_p_user@gmail.com', '', 23, NULL, '2023-01-05 17:10:37', '2023-01-05 17:10:37'),
(67, 'user_i_e', 'p_q_user@gmail.com', '', 23, NULL, '2023-01-05 17:10:38', '2023-01-05 17:10:38'),
(68, 'user_x_d', 's_h_user@gmail.com', '', 23, NULL, '2023-01-05 17:10:39', '2023-01-05 17:10:39');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `movies`
--
ALTER TABLE `movies`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `Code_UniqueIndex` (`moviecode`);

--
-- Indexes for table `person`
--
ALTER TABLE `person`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `emailIndexUnique` (`email`) USING BTREE;

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `movies`
--
ALTER TABLE `movies`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `person`
--
ALTER TABLE `person`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=69;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
