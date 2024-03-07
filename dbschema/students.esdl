module Students {
    type Students {
        required name: str;
        school: str;
        class: str;
        required email: str;
        required created_at: datetime {
            readonly := true;
            default := datetime_of_statement();
        };
        updated_at: datetime {
            rewrite update using (datetime_of_statement());
        };
    }
}