
/**
 * Samples for CollectionPartition ListUsages.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBCollectionPartitionGetUsages.json
     */
    /**
     * Sample code: CosmosDBCollectionGetUsages.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCollectionGetUsages(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCollectionPartitions().listUsages("rg1", "ddb1", "databaseRid", "collectionRid",
            "name.value eq 'Partition Storage'", com.azure.core.util.Context.NONE);
    }
}
