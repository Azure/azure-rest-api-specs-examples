
/**
 * Samples for Servers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/ServerDelete
     * .json
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
