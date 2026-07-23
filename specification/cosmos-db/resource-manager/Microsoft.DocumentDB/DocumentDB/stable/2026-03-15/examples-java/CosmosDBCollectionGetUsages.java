
/**
 * Samples for Collection ListUsages.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBCollectionGetUsages.json
     */
    /**
     * Sample code: CosmosDBCollectionGetUsages.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCollectionGetUsages(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCollections().listUsages("rg1", "ddb1", "databaseRid", "collectionRid",
            "name.value eq 'Storage'", com.azure.core.util.Context.NONE);
    }
}
