Do to issues with the export I could not make an export thats not plain text, the database can be recreated by using the following commands:

- Make a Database with the name: 'password_generator'
- Run the following in the database: 
 1).   
    CREATE TABLE `passwords` (
    `id` int(11) NOT NULL,
    `Password` varchar(502) NOT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

----------------------------------------------------------------------------------------------------------------------------------------
2).
    INSERT INTO `passwords` (`id`, `Password`) VALUES
    (0, 'MrYb'),
    (16, 'LcdYwUYuUWsUUIpf');

----------------------------------------------------------------------------------------------------------------------------------------
3).
    ALTER TABLE `passwords`
        ADD PRIMARY KEY (`id`);

----------------------------------------------------------------------------------------------------------------------------------------
4).
    ALTER TABLE `passwords`
        MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;
    COMMIT;

----------------------------------------------------------------------------------------------------------------------------------------

Afterwards change the DB credentails on line 15 - 17. 