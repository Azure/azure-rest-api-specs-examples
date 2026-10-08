
/**
 * Samples for CopyJobs Complete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/copy-jobs/CosmosDBCopyJobComplete.json
     */
    /**
     * Sample code: CosmosDBCopyJobComplete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCopyJobComplete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCopyJobs().completeWithResponse("rg1", "ddb1", "j1",
            com.azure.core.util.Context.NONE);
    }
}
