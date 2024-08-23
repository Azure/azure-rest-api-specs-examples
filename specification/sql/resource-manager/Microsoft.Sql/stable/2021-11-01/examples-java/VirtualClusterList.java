
/**
 * Samples for VirtualClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/VirtualClusterList.json
     */
    /**
     * Sample code: List virtualClusters.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVirtualClusters(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getVirtualClusters().list(com.azure.core.util.Context.NONE);
    }
}
