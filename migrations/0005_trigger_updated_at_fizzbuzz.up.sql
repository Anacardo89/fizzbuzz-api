CREATE TRIGGER fizzbuzz_set_updated_at
BEFORE UPDATE ON fizzbuzz
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
