
/**
 * Samples for CopyJobs Resume.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/copy-jobs/CosmosDBCopyJobResume.json
     */
    /**
     * Sample code: CosmosDBCopyJobResume.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCopyJobResume(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCopyJobs().resumeWithResponse("rg1", "ddb1", "j1", com.azure.core.util.Context.NONE);
    }
}
