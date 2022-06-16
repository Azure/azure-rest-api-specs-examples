import com.azure.core.util.Context;

/** Samples for Databases ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/DatabaseListByServer.json
     */
    /**
     * Sample code: DatabaseList.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void databaseList(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.databases().listByServer("TestGroup", "testserver", Context.NONE);
    }
}
