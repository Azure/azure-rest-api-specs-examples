
/**
 * Samples for ServerTrustGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerTrustGroupGet.json
     */
    /**
     * Sample code: Get server trust group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getServerTrustGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerTrustGroups().getWithResponse("Default", "Japan East",
            "server-trust-group-test", com.azure.core.util.Context.NONE);
    }
}
