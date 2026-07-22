
/**
 * Samples for PartitionKeyRangeId ListMetrics.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBPKeyRangeIdGetMetrics.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountRegionGetMetrics.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDatabaseAccountRegionGetMetrics(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getPartitionKeyRangeIds().listMetrics("rg1", "ddb1", "databaseRid", "collectionRid",
            "0",
            "(name.value eq 'Max RUs Per Second') and timeGrain eq duration'PT1M' and startTime eq '2017-11-19T23:53:55.2780000Z' and endTime eq '2017-11-20T23:58:55.2780000Z",
            com.azure.core.util.Context.NONE);
    }
}
