CREATE TABLE `mvs` (
  `tmdb` int PRIMARY KEY,
  `title` varchar(255),
  `tagline` varchar(255),
  `release` string,
  `runtime` int,
  `genres` varchar(255),
  `overview` text,
  `poster` varchar(255),
  `backdrop` varchar(255),
  `trailers` varchar(255)
);
