
/**
 * Samples for Clusters ListNamespaces.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/Clusters/
     * ListNamespacesInClusterGet.json
     */
    /**
     * Sample code: ListNamespacesInCluster.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listNamespacesInCluster(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getClusters().listNamespacesWithResponse("myResourceGroup",
            "testCluster", com.azure.core.util.Context.NONE);
    }
}
