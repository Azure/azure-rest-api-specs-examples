
/**
 * Samples for ServerDevOpsAuditSettings ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerDevOpsAuditList.json
     */
    /**
     * Sample code: List DevOps audit settings of a server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listDevOpsAuditSettingsOfAServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerDevOpsAuditSettings().listByServer("devAuditTestRG", "devOpsAuditTestSvr",
            com.azure.core.util.Context.NONE);
    }
}
