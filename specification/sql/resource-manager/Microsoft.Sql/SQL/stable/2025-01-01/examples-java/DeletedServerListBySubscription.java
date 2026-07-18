
/**
 * Samples for DeletedServers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DeletedServerListBySubscription.json
     */
    /**
     * Sample code: List deleted servers in a subscription.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listDeletedServersInASubscription(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDeletedServers().list(com.azure.core.util.Context.NONE);
    }
}
