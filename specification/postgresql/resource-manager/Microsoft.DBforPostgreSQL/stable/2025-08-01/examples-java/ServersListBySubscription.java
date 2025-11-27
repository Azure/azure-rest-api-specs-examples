
/**
 * Samples for Servers List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * ServersListBySubscription.json
     */
    /**
     * Sample code: List all servers in a subscription.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        listAllServersInASubscription(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().list(com.azure.core.util.Context.NONE);
    }
}
