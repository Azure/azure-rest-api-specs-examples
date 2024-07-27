
/**
 * Samples for Clusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/Clusters/
     * ClustersListByResourceGroup.json
     */
    /**
     * Sample code: ClustersListByResourceGroup.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void clustersListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getClusters().listByResourceGroup("myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
