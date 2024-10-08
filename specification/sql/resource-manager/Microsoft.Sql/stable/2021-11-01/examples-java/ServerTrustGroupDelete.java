
/**
 * Samples for ServerTrustGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerTrustGroupDelete.json
     */
    /**
     * Sample code: Drop server trust group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dropServerTrustGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerTrustGroups().delete("Default", "Japan East",
            "server-trust-group-test", com.azure.core.util.Context.NONE);
    }
}
