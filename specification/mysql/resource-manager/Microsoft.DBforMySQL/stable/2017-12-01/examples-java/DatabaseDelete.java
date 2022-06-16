import com.azure.core.util.Context;

/** Samples for Databases Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/DatabaseDelete.json
     */
    /**
     * Sample code: DatabaseDelete.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void databaseDelete(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.databases().delete("TestGroup", "testserver", "db1", Context.NONE);
    }
}
