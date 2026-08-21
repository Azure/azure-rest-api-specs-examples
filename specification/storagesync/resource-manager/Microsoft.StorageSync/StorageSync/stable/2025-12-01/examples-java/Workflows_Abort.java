
/**
 * Samples for Workflows Abort.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/Workflows_Abort.json
     */
    /**
     * Sample code: Workflows_Abort.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void workflowsAbort(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.workflows().abortWithResponse("SampleResourceGroup_1", "SampleStorageSyncService_1",
            "7ffd50b3-5574-478d-9ff2-9371bc42ce68", com.azure.core.util.Context.NONE);
    }
}
