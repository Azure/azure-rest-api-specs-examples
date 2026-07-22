
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBOperationsList.json
     */
    /**
     * Sample code: CosmosDBOperationsList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBOperationsList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
