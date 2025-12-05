
/**
 * Samples for NetAppResourceUsages Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Usages_Get.json
     */
    /**
     * Sample code: Usages_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void usagesGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.netAppResourceUsages().getWithResponse("eastus", "totalTibsPerSubscription",
            com.azure.core.util.Context.NONE);
    }
}
