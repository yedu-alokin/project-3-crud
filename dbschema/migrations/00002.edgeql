CREATE MIGRATION m177yrk4go4bq4mxcjwrenzn4gig4v5m6uismng2sssgl3j4rmmqya
    ONTO m1pkeqsgl2udqce6xtmrqm757q4t3y7fjbanpmfko2osccsmjuwgha
{
  ALTER TYPE Students::Students {
      ALTER PROPERTY email {
          DROP CONSTRAINT std::exclusive;
      };
  };
};
