-- SQL Dump
-- version 1.0.0
-- https://www.edwinndui.com/
--
-- Host: 127.0.0.1
-- Generation Time: 16 Jan 2022 15:12
-- Server version: 10.3.16-MariaDB

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `houseinfodb`
--

-- --------------------------------------------------------

--
-- Table structure for table `house`
--

CREATE TABLE `house` (
  `id` int(11) NOT NULL,
  `house_number` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `location` varchar(100) NOT NULL,
  `cost` int(11) NOT NULL,
  `year_of_construction` int(11) NOT NULL,
  `category` varchar(100) NOT NULL,
  `area` float(11) NOT NULL,
  `perimeter` float(11) NOT NULL,
  `number_of_floors` int(11) NOT NULL,
  `number_of_bedrooms` int(11) NOT NULL,
  `construction_material` varchar(100) NOT NULL,
  `roofing_type` varchar(100) NOT NULL,
  `fencing_type` varchar(100) NOT NULL,
  `parking_lot` TINYINT(1) DEFAULT '0',
  `source_of_water_supply` varchar(100) NOT NULL,
  `has_wifi` TINYINT(1) DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


--
-- Indexes for dumped tables
--

--
-- Indexes for table `house`
--
ALTER TABLE `house`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `house`
--
ALTER TABLE `house`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
