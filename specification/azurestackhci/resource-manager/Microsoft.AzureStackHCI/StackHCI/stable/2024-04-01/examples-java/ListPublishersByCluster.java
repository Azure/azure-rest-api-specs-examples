
/**
 * Samples for Publishers ListByCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * ListPublishersByCluster.json
     */
    /**
     * Sample code: List Publisher resources by HCI Cluster.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        listPublisherResourcesByHCICluster(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.publishers().listByCluster("test-rg", "myCluster", com.azure.core.util.Context.NONE);
    }
}
