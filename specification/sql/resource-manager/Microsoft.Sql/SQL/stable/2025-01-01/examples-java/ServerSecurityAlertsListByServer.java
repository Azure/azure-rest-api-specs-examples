
/**
 * Samples for ServerSecurityAlertPolicies ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerSecurityAlertsListByServer.json
     */
    /**
     * Sample code: List the server's threat detection policies.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listTheServerSThreatDetectionPolicies(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerSecurityAlertPolicies().listByServer("securityalert-4799",
            "securityalert-6440", com.azure.core.util.Context.NONE);
    }
}
