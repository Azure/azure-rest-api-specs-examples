
/**
 * Samples for VolumeQuotaRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/VolumeQuotaRules_Delete.json
     */
    /**
     * Sample code: VolumeQuotaRules_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeQuotaRulesDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumeQuotaRules().delete("myRG", "account-9957", "pool-5210", "volume-6387", "rule-0004",
            com.azure.core.util.Context.NONE);
    }
}
