
/**
 * Samples for Servers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ServersListByResourceGroup.json
     */
    /**
     * Sample code: List all servers in a resource group.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        listAllServersInAResourceGroup(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().listByResourceGroup("exampleresourcegroup", com.azure.core.util.Context.NONE);
    }
}
