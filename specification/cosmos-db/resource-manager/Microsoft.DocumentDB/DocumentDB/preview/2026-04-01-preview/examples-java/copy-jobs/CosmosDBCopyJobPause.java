
/**
 * Samples for CopyJobs Pause.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/copy-jobs/CosmosDBCopyJobPause.json
     */
    /**
     * Sample code: CosmosDBCopyJobPause.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCopyJobPause(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCopyJobs().pauseWithResponse("rg1", "ddb1", "j1", com.azure.core.util.Context.NONE);
    }
}
