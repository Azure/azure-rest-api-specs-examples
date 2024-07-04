
/**
 * Samples for DnsPrivateViews Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dnsPrivateViews_get.json
     */
    /**
     * Sample code: Get a DnsPrivateView by name.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        getADnsPrivateViewByName(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dnsPrivateViews().getWithResponse("eastus", "ocid1....aaaaaa", com.azure.core.util.Context.NONE);
    }
}
