
/**
 * Samples for Clusters List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * ListClustersBySubscription.json
     */
    /**
     * Sample code: List clusters in a given subscription.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        listClustersInAGivenSubscription(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusters().list(com.azure.core.util.Context.NONE);
    }
}
