/** Samples for Databases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/DatabaseCreate.json
     */
    /**
     * Sample code: DatabaseCreate.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void databaseCreate(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager
            .databases()
            .define("db1")
            .withExistingServer("TestGroup", "testserver")
            .withCharset("UTF8")
            .withCollation("English_United States.1252")
            .create();
    }
}
