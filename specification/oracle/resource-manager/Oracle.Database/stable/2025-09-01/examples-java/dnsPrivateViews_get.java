
/**
 * Samples for DnsPrivateViews Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/dnsPrivateViews_get.json
     */
    /**
     * Sample code: DnsPrivateViews_get.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void dnsPrivateViewsGet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dnsPrivateViews().getWithResponse("eastus", "ocid1....aaaaaa", com.azure.core.util.Context.NONE);
    }
}
