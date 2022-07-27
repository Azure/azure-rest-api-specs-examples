import com.azure.core.util.Context;

/** Samples for VirtualClusters ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2015-05-01-preview/examples/VirtualClusterListByResourceGroup.json
     */
    /**
     * Sample code: List virtual clusters by resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVirtualClustersByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getVirtualClusters().listByResourceGroup("testrg", Context.NONE);
    }
}
