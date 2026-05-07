
/**
 * Samples for NetAppResourceQuotaLimitsAccount Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/NetAppResourceQuotaLimitsAccount_Get.json
     */
    /**
     * Sample code: QuotaLimits.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void quotaLimits(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.netAppResourceQuotaLimitsAccounts().getWithResponse("myRG", "myAccount", "poolsPerAccount",
            com.azure.core.util.Context.NONE);
    }
}
