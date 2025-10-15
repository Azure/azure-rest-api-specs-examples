
/**
 * Samples for DnsPrivateViews ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/dnsPrivateViews_listByLocation.json
     */
    /**
     * Sample code: DnsPrivateViews_listByLocation.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        dnsPrivateViewsListByLocation(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dnsPrivateViews().listByLocation("eastus", com.azure.core.util.Context.NONE);
    }
}
