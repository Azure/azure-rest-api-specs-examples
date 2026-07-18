
/**
 * Samples for ExtendedServerBlobAuditingPolicies ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerExtendedAuditingSettingsList.json
     */
    /**
     * Sample code: List extended auditing settings of a server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listExtendedAuditingSettingsOfAServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getExtendedServerBlobAuditingPolicies().listByServer("blobauditingtest-4799",
            "blobauditingtest-6440", com.azure.core.util.Context.NONE);
    }
}
