
/**
 * Samples for NeuroShieldOperationStatuses Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldOperationStatuses_Get.json
     */
    /**
     * Sample code: NeuroShieldOperationStatuses_Get.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void
        neuroShieldOperationStatusesGet(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldOperationStatuses().getWithResponse("eastus2", "00000000-0000-0000-0000-000000000001",
            com.azure.core.util.Context.NONE);
    }
}
