
/**
 * Samples for Servers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * ServerListByResourceGroup.json
     */
    /**
     * Sample code: ServerListByResourceGroup.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        serverListByResourceGroup(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().listByResourceGroup("testrgn", com.azure.core.util.Context.NONE);
    }
}
