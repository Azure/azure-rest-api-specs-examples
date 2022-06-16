import com.azure.core.util.Context;

/** Samples for Servers ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerListByResourceGroup.json
     */
    /**
     * Sample code: ServerListByResourceGroup.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void serverListByResourceGroup(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.servers().listByResourceGroup("testrg", Context.NONE);
    }
}
