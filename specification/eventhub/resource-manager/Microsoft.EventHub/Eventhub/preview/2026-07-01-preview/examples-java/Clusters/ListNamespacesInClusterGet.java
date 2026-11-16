
/**
 * Samples for Clusters ListNamespaces.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Clusters/ListNamespacesInClusterGet.json
     */
    /**
     * Sample code: ListNamespacesInCluster.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void listNamespacesInCluster(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getClusters().listNamespacesWithResponse("myResourceGroup", "testCluster",
            com.azure.core.util.Context.NONE);
    }
}
