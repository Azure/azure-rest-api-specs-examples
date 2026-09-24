
/**
 * Samples for ServerTrustGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ServerTrustGroupDelete.json
     */
    /**
     * Sample code: Drop server trust group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void dropServerTrustGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerTrustGroups().delete("Default", "Japan East", "server-trust-group-test",
            com.azure.core.util.Context.NONE);
    }
}
