
/**
 * Samples for VolumeQuotaRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/VolumeQuotaRules_Get.json
     */
    /**
     * Sample code: VolumeQuotaRules_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeQuotaRulesGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumeQuotaRules().getWithResponse("myRG", "account-9957", "pool-5210", "volume-6387", "rule-0004",
            com.azure.core.util.Context.NONE);
    }
}
