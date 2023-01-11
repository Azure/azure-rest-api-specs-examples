/** Samples for Databases Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/examples/DatabaseGet.json
     */
    /**
     * Sample code: Get a database.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getADatabase(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.databases().getWithResponse("TestGroup", "testserver", "db1", com.azure.core.util.Context.NONE);
    }
}
