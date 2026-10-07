
/**
 * Samples for GarnetClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBGarnetClusterListByResourceGroup.json
     */
    /**
     * Sample code: CosmosDBGarnetClusterListByResourceGroup.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBGarnetClusterListByResourceGroup(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGarnetClusters().listByResourceGroup("garnet-prod-rg",
            com.azure.core.util.Context.NONE);
    }
}
