
/**
 * Samples for Fleet GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/fleet/CosmosDBFleetGet.json
     */
    /**
     * Sample code: CosmosDB Fleet Get.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBFleetGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getFleets().getByResourceGroupWithResponse("rg1", "fleet1",
            com.azure.core.util.Context.NONE);
    }
}
