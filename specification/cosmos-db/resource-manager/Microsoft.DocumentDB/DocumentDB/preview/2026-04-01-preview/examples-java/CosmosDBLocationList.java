
/**
 * Samples for Locations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBLocationList.json
     */
    /**
     * Sample code: CosmosDBLocationList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBLocationList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getLocations().list(com.azure.core.util.Context.NONE);
    }
}
