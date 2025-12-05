
/**
 * Samples for NetAppResourceQuotaLimitsAccount List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/NetAppResourceQuotaLimitsAccount_List.json
     */
    /**
     * Sample code: QuotaLimits.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void quotaLimits(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.netAppResourceQuotaLimitsAccounts().list("myRG", "myAccount", com.azure.core.util.Context.NONE);
    }
}
