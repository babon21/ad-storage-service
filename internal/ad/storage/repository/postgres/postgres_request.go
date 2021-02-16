package postgres

const createAdQuery = "INSERT INTO ad(title,description,price,main_photo,photos,date_added) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id"

const getSimpleAdQuery = "SELECT title,price,main_photo FROM ad WHERE id = $1"

const getAdWithPhotosQuery = "SELECT title,price,main_photo,photos FROM ad WHERE id = $1"

const getAdWithDescQuery = "SELECT title,price,description,main_photo FROM ad WHERE id = $1"

const getAdWithPhotoAndDescQuery = "SELECT title,price,description,main_photo,photos FROM ad WHERE id = $1"
