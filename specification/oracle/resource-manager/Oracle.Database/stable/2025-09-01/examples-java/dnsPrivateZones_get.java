
/**
 * Samples for DnsPrivateZones Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/dnsPrivateZones_get.json
     */
    /**
     * Sample code: DnsPrivateZones_get.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void dnsPrivateZonesGet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dnsPrivateZones().getWithResponse("eastus", "example-dns-private-zone",
            com.azure.core.util.Context.NONE);
    }
}
