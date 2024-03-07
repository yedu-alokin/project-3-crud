CREATE MIGRATION m1pkeqsgl2udqce6xtmrqm757q4t3y7fjbanpmfko2osccsmjuwgha
    ONTO initial
{
  CREATE MODULE Students IF NOT EXISTS;
  CREATE TYPE Students::Students {
      CREATE PROPERTY class: std::str;
      CREATE REQUIRED PROPERTY created_at: std::datetime {
          SET default := (std::datetime_of_statement());
          SET readonly := true;
      };
      CREATE REQUIRED PROPERTY email: std::str {
          CREATE CONSTRAINT std::exclusive;
      };
      CREATE REQUIRED PROPERTY name: std::str;
      CREATE PROPERTY school: std::str;
      CREATE PROPERTY updated_at: std::datetime {
          CREATE REWRITE
              UPDATE 
              USING (std::datetime_of_statement());
      };
  };
};
