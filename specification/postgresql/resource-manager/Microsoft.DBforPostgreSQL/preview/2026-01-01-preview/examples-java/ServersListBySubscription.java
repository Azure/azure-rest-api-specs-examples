
/**
 * Samples for Servers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ServersListBySubscription.json
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
