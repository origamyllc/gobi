package sql

const GET_USERS_QUERY  = `SELECT * FROM users`;
const GET_USER_BY_ID_QUERY = "SELECT * FROM users WHERE ID =";
const ADD_USER_QUERY = "INSERT INTO users (age, email, first_name, last_name) VALUES (%d,'%s','%s','%s');"
const DELETE_USER_BY_ID_QUERY = "DELETE FROM users WHERE id = "
const UPDATE_USER_NAME__QUERY =  `UPDATE users SET first_name = '%s', last_name = '%s' WHERE id = %d;`