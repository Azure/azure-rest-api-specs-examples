
/**
 * Samples for NetAppResourceQuotaLimits Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/NetAppResourceQuotaLimits_Get.json
     */
    /**
     * Sample code: QuotaLimits.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void quotaLimits(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.netAppResourceQuotaLimits().getWithResponse("eastus", "totalCoolAccessVolumesPerSubscription",
            com.azure.core.util.Context.NONE);
    }
}
