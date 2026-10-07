
/**
 * Samples for CopyJobs ListByDatabaseAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/copy-jobs/CosmosDBCopyJobFeed.json
     */
    /**
     * Sample code: CosmosDBCopyJobFeed.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCopyJobFeed(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCopyJobs().listByDatabaseAccount("rg1", "ddb1", com.azure.core.util.Context.NONE);
    }
}
