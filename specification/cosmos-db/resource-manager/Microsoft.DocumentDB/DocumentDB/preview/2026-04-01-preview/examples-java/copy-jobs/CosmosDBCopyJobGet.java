
/**
 * Samples for CopyJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/copy-jobs/CosmosDBCopyJobGet.json
     */
    /**
     * Sample code: CosmosDBCopyJobGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCopyJobGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCopyJobs().getWithResponse("rg1", "ddb1", "j1", com.azure.core.util.Context.NONE);
    }
}
