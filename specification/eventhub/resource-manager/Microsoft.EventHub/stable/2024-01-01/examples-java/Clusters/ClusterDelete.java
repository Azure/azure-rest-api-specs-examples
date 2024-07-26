
/**
 * Samples for Clusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/Clusters/ClusterDelete.json
     */
    /**
     * Sample code: ClusterDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void clusterDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getClusters().delete("myResourceGroup", "testCluster",
            com.azure.core.util.Context.NONE);
    }
}
