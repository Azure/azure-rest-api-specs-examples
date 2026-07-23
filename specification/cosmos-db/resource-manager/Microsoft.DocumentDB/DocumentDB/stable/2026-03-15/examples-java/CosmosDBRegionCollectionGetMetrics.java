
/**
 * Samples for CollectionRegion ListMetrics.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBRegionCollectionGetMetrics.json
     */
    /**
     * Sample code: CosmosDBRegionCollectionGetMetrics.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBRegionCollectionGetMetrics(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCollectionRegions().listMetrics("rg1", "ddb1", "North Europe", "databaseRid",
            "collectionRid",
            "(name.value eq 'Total Requests') and timeGrain eq duration'PT5M' and startTime eq '2017-11-19T23:53:55.2780000Z' and endTime eq '2017-11-20T00:13:55.2780000Z",
            com.azure.core.util.Context.NONE);
    }
}
