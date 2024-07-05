
/**
 * Samples for DnsPrivateZones Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dnsPrivateZones_get.json
     */
    /**
     * Sample code: Get a DnsPrivateZone by name.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        getADnsPrivateZoneByName(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dnsPrivateZones().getWithResponse("eastus", "example-dns-private-zone",
            com.azure.core.util.Context.NONE);
    }
}
