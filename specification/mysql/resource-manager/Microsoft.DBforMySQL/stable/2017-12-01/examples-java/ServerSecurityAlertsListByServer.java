import com.azure.core.util.Context;

/** Samples for ServerSecurityAlertPolicies ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerSecurityAlertsListByServer.json
     */
    /**
     * Sample code: List the server's threat detection policies.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void listTheServerSThreatDetectionPolicies(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.serverSecurityAlertPolicies().listByServer("securityalert-4799", "securityalert-6440", Context.NONE);
    }
}
