
/**
 * Samples for Clusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * ListClustersByResourceGroup.json
     */
    /**
     * Sample code: List clusters in a given resource group.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        listClustersInAGivenResourceGroup(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusters().listByResourceGroup("test-rg", com.azure.core.util.Context.NONE);
    }
}
