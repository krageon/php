<?php
require "connect.inc";

$link = ldap_connect_and_bind($host, $port, $user, $passwd, $protocol_version);
insert_dummy_data($link, $base);
$result = ldap_search($link, "$base", "(o=test)");
$entry = ldap_first_entry($link, $result);
var_dump(
	ldap_get_attributes($link, $entry)
);
?>
===DONE===
