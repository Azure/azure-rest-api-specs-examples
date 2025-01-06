
/**
 * Samples for Servers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/
     * ServerListByResourceGroup.json
     */
    /**
     * Sample code: ServerListByResourceGroup.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        serverListByResourceGroup(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().listByResourceGroup("testrg", com.azure.core.util.Context.NONE);
    }
}
