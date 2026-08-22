
/**
 * Samples for Workflows Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/Workflows_Get.json
     */
    /**
     * Sample code: Workflows_Get.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void workflowsGet(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.workflows().getWithResponse("SampleResourceGroup_1", "SampleStorageSyncService_1",
            "828219ea-083e-48b5-89ea-8fd9991b2e75", com.azure.core.util.Context.NONE);
    }
}
