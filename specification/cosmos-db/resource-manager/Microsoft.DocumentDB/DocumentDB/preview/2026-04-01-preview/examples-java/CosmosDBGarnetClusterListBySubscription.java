
/**
 * Samples for GarnetClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBGarnetClusterListBySubscription.json
     */
    /**
     * Sample code: CosmosDBGarnetClusterListBySubscription.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGarnetClusterListBySubscription(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGarnetClusters().list(com.azure.core.util.Context.NONE);
    }
}
