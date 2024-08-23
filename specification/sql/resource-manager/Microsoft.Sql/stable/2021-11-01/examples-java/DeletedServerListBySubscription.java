
/**
 * Samples for DeletedServers List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DeletedServerListBySubscription.json
     */
    /**
     * Sample code: List deleted servers in a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDeletedServersInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDeletedServers().list(com.azure.core.util.Context.NONE);
    }
}
