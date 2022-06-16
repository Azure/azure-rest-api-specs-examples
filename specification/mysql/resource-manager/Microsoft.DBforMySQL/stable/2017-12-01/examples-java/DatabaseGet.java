import com.azure.core.util.Context;

/** Samples for Databases Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/DatabaseGet.json
     */
    /**
     * Sample code: DatabaseGet.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void databaseGet(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.databases().getWithResponse("TestGroup", "testserver", "db1", Context.NONE);
    }
}
