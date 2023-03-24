import com.azure.core.util.Context;

/** Samples for Databases Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/DatabaseDelete.json
     */
    /**
     * Sample code: Delete a database.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void deleteADatabase(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.databases().delete("TestGroup", "testserver", "db1", Context.NONE);
    }
}
