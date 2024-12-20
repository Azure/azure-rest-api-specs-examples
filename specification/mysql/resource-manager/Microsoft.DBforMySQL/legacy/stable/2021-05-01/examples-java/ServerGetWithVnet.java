
/**
 * Samples for Servers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2021-05-01/examples/ServerGetWithVnet.
     * json
     */
    /**
     * Sample code: Get a server with vnet.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void getAServerWithVnet(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.servers().getByResourceGroupWithResponse("testrg", "mysqltestserver", com.azure.core.util.Context.NONE);
    }
}
