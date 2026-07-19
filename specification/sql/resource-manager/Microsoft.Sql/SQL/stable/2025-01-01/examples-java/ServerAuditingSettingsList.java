
/**
 * Samples for ServerBlobAuditingPolicies ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerAuditingSettingsList.json
     */
    /**
     * Sample code: List auditing settings of a server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listAuditingSettingsOfAServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerBlobAuditingPolicies().listByServer("blobauditingtest-4799",
            "blobauditingtest-6440", com.azure.core.util.Context.NONE);
    }
}
