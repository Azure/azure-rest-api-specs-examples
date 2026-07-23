
/**
 * Samples for RestorableMongodbCollections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBRestorableMongodbCollectionList.json
     */
    /**
     * Sample code: CosmosDBRestorableMongodbCollectionList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBRestorableMongodbCollectionList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getRestorableMongodbCollections().list("WestUS", "98a570f2-63db-4117-91f0-366327b7b353",
            "PD5DALigDgw=", null, null, com.azure.core.util.Context.NONE);
    }
}
