
/**
 * Samples for CopyJobs Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/copy-jobs/CosmosDBCopyJobCancel.json
     */
    /**
     * Sample code: CosmosDBCopyJobCancel.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCopyJobCancel(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCopyJobs().cancelWithResponse("rg1", "ddb1", "j1", com.azure.core.util.Context.NONE);
    }
}
