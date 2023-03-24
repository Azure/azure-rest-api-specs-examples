import com.azure.core.util.Context;

/** Samples for Databases Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/DatabaseGet.json
     */
    /**
     * Sample code: Get a database.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void getADatabase(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.databases().getWithResponse("TestGroup", "testserver", "db1", Context.NONE);
    }
}
