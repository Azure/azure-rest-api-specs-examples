
/**
 * Samples for DnsPrivateZones ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/dnsPrivateZones_listByLocation.json
     */
    /**
     * Sample code: DnsPrivateZones_listByLocation.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        dnsPrivateZonesListByLocation(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dnsPrivateZones().listByLocation("eastus", com.azure.core.util.Context.NONE);
    }
}
