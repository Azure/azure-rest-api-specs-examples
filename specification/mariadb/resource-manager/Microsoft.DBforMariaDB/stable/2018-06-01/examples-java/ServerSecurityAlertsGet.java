import com.azure.resourcemanager.mariadb.models.SecurityAlertPolicyName;

/** Samples for ServerSecurityAlertPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerSecurityAlertsGet.json
     */
    /**
     * Sample code: Get a server's threat detection policy.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void getAServerSThreatDetectionPolicy(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager
            .serverSecurityAlertPolicies()
            .getWithResponse(
                "securityalert-4799",
                "securityalert-6440",
                SecurityAlertPolicyName.DEFAULT,
                com.azure.core.util.Context.NONE);
    }
}
