
/**
 * Samples for Clusters List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/Clusters/
     * ClustersListBySubscription.json
     */
    /**
     * Sample code: ClustersListBySubscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void clustersListBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getClusters().list(com.azure.core.util.Context.NONE);
    }
}
