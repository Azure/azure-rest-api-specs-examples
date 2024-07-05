
/**
 * Samples for DnsPrivateViews ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dnsPrivateViews_listByLocation.
     * json
     */
    /**
     * Sample code: List DnsPrivateViews by location.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        listDnsPrivateViewsByLocation(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dnsPrivateViews().listByLocation("eastus", com.azure.core.util.Context.NONE);
    }
}
