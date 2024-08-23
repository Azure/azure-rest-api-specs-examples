
/**
 * Samples for VirtualClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/VirtualClusterListByResourceGroup.
     * json
     */
    /**
     * Sample code: List virtual clusters by resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVirtualClustersByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getVirtualClusters().listByResourceGroup("testrg",
            com.azure.core.util.Context.NONE);
    }
}
