import com.azure.core.util.Context;

/** Samples for ServerSecurityAlertPolicies ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerSecurityAlertsListByServer.json
     */
    /**
     * Sample code: List the server's threat detection policies.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listTheServerSThreatDetectionPolicies(
        com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.serverSecurityAlertPolicies().listByServer("securityalert-4799", "securityalert-6440", Context.NONE);
    }
}
