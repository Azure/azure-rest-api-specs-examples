
/**
 * Samples for NetAppResourceQuotaLimits Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-01-01-preview/examples/QuotaLimits_Get.json
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
