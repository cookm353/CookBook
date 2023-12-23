.echo ON

DROP TABLE IF EXISTS 'recipe_ingredients';
CREATE TABLE 'recipe_ingredients' (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    recipe_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    amount TEXT NOT NULL,
    unit TEXT NOT NULL,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id)
);

DROP TABLE IF EXISTS 'recipes';
CREATE TABLE 'recipes' (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL
);