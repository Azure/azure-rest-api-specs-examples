
/**
 * Samples for Locations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBLocationGet.json
     */
    /**
     * Sample code: CosmosDBLocationGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBLocationGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getLocations().getWithResponse("westus", com.azure.core.util.Context.NONE);
    }
}
