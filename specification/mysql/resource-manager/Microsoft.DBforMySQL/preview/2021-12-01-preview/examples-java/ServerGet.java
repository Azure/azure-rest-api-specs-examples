/** Samples for Servers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/preview/2021-12-01-preview/examples/ServerGet.json
     */
    /**
     * Sample code: Get a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void getAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.servers().getByResourceGroupWithResponse("testrg", "mysqltestserver", com.azure.core.util.Context.NONE);
    }
}
