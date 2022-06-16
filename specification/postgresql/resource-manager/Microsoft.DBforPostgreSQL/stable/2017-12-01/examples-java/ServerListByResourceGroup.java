import com.azure.core.util.Context;

/** Samples for Servers ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerListByResourceGroup.json
     */
    /**
     * Sample code: ServerListByResourceGroup.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverListByResourceGroup(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.servers().listByResourceGroup("TestGroup", Context.NONE);
    }
}
