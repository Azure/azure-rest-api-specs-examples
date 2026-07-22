
/**
 * Samples for CollectionPartitionRegion ListMetrics.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBCollectionPartitionRegionGetMetrics.json
     */
    /**
     * Sample code: CosmosDBCollectionPartitionRegionGetMetrics.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBCollectionPartitionRegionGetMetrics(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCollectionPartitionRegions().listMetrics("rg1", "ddb1", "North Europe",
            "databaseRid", "collectionRid",
            "(name.value eq 'Max RUs Per Second') and timeGrain eq duration'PT1M' and startTime eq '2017-11-19T23:53:55.2780000Z' and endTime eq '2017-11-20T23:58:55.2780000Z",
            com.azure.core.util.Context.NONE);
    }
}
