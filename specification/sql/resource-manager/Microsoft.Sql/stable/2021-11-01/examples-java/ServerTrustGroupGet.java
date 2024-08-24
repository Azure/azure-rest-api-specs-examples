
/**
 * Samples for ServerTrustGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerTrustGroupGet.json
     */
    /**
     * Sample code: Get server trust group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getServerTrustGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerTrustGroups().getWithResponse("Default", "Japan East",
            "server-trust-group-test", com.azure.core.util.Context.NONE);
    }
}
