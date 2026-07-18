
/**
 * Samples for ServerTrustGroups ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerTrustGroupList.json
     */
    /**
     * Sample code: List server trust groups.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listServerTrustGroups(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerTrustGroups().listByLocation("Default", "Japan East",
            com.azure.core.util.Context.NONE);
    }
}
