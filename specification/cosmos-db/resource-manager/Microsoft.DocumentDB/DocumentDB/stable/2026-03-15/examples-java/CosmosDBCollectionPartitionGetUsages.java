
/**
 * Samples for CollectionPartition ListUsages.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBCollectionPartitionGetUsages.json
     */
    /**
     * Sample code: CosmosDBCollectionPartitionGetUsages.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCollectionPartitionGetUsages(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCollectionPartitions().listUsages("rg1", "ddb1", "databaseRid", "collectionRid",
            "name.value eq 'Partition Storage'", com.azure.core.util.Context.NONE);
    }
}
