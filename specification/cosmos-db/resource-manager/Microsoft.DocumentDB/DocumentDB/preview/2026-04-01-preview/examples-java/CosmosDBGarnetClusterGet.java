
/**
 * Samples for GarnetClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBGarnetClusterGet.json
     */
    /**
     * Sample code: CosmosDBGarnetClusterGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGarnetClusterGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGarnetClusters().getByResourceGroupWithResponse("garnet-prod-rg", "garnet-prod",
            com.azure.core.util.Context.NONE);
    }
}
