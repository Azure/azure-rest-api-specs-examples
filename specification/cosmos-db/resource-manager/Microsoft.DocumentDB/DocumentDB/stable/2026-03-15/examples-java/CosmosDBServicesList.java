
/**
 * Samples for Service List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBServicesList.json
     */
    /**
     * Sample code: CosmosDBServicesList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBServicesList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getServices().list("rg1", "ddb1", com.azure.core.util.Context.NONE);
    }
}
