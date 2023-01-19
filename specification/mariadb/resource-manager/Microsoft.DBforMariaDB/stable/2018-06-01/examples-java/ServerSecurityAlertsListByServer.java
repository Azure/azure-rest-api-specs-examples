/** Samples for ServerSecurityAlertPolicies ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerSecurityAlertsListByServer.json
     */
    /**
     * Sample code: List the server's threat detection policies.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void listTheServerSThreatDetectionPolicies(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager
            .serverSecurityAlertPolicies()
            .listByServer("securityalert-4799", "securityalert-6440", com.azure.core.util.Context.NONE);
    }
}
