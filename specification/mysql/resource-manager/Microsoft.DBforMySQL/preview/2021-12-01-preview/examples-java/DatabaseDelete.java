/** Samples for Databases Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/preview/2021-12-01-preview/examples/DatabaseDelete.json
     */
    /**
     * Sample code: Delete a database.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void deleteADatabase(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.databases().delete("TestGroup", "testserver", "db1", com.azure.core.util.Context.NONE);
    }
}
