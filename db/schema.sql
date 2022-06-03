
CREATE TABLE IF NOT EXISTS url (
                                   id INTEGER PRIMARY KEY AUTOINCREMENT,
                                   full_url TEXT NOT NULL,
                                   short_url TEXT NOT NULL UNIQUE,
                                   counter INTEGER NOT NULL
);


