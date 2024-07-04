
/**
 * Samples for DnsPrivateZones ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dnsPrivateZones_listByLocation.
     * json
     */
    /**
     * Sample code: List DnsPrivateZones by location.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        listDnsPrivateZonesByLocation(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dnsPrivateZones().listByLocation("eastus", com.azure.core.util.Context.NONE);
    }
}
