import com.azure.core.util.Context;
import com.azure.resourcemanager.mysql.models.SecurityAlertPolicyName;

/** Samples for ServerSecurityAlertPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerSecurityAlertsGet.json
     */
    /**
     * Sample code: Get a server's threat detection policy.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void getAServerSThreatDetectionPolicy(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .serverSecurityAlertPolicies()
            .getWithResponse("securityalert-4799", "securityalert-6440", SecurityAlertPolicyName.DEFAULT, Context.NONE);
    }
}
