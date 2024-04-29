
/**
 * Samples for Servers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/
     * ServersListByResourceGroup.json
     */
    /**
     * Sample code: List servers in a resource group.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void listServersInAResourceGroup(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.servers().listByResourceGroup("TestGroup", com.azure.core.util.Context.NONE);
    }
}
