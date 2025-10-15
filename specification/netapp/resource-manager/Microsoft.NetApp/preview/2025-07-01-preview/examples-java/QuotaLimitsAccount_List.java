
/**
 * Samples for NetAppResourceQuotaLimitsAccount List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-07-01-preview/examples/
     * QuotaLimitsAccount_List.json
     */
    /**
     * Sample code: QuotaLimitsAccount_List.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void quotaLimitsAccountList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.netAppResourceQuotaLimitsAccounts().list("myRG", "myAccount", com.azure.core.util.Context.NONE);
    }
}
