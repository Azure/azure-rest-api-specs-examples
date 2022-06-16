import com.azure.core.util.Context;

/** Samples for Servers ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ServersListByResourceGroup.json
     */
    /**
     * Sample code: List servers in a resource group.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void listServersInAResourceGroup(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.servers().listByResourceGroup("TestGroup", Context.NONE);
    }
}
