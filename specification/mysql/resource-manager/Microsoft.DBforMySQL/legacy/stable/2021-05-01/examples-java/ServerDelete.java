
/**
 * Samples for Servers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2021-05-01/examples/ServerDelete.json
     */
    /**
     * Sample code: Delete a server.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void deleteAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.servers().delete("TestGroup", "testserver", com.azure.core.util.Context.NONE);
    }
}
