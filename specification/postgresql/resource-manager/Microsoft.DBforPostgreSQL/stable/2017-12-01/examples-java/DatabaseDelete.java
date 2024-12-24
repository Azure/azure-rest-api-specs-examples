
/**
 * Samples for Databases Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/DatabaseDelete.
     * json
     */
    /**
     * Sample code: DatabaseDelete.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void databaseDelete(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.databases().delete("TestGroup", "testserver", "db1", com.azure.core.util.Context.NONE);
    }
}
