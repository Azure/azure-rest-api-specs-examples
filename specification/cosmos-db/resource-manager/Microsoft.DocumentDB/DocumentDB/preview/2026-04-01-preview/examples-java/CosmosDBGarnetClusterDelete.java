
/**
 * Samples for GarnetClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBGarnetClusterDelete.json
     */
    /**
     * Sample code: CosmosDBGarnetClusterDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGarnetClusterDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGarnetClusters().delete("garnet-prod-rg", "garnet-prod",
            com.azure.core.util.Context.NONE);
    }
}
