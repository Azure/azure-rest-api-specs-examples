import com.azure.core.util.Context;
import com.azure.resourcemanager.postgresql.models.SecurityAlertPolicyName;

/** Samples for ServerSecurityAlertPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerSecurityAlertsGet.json
     */
    /**
     * Sample code: Get a server's threat detection policy.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getAServerSThreatDetectionPolicy(
        com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager
            .serverSecurityAlertPolicies()
            .getWithResponse("securityalert-4799", "securityalert-6440", SecurityAlertPolicyName.DEFAULT, Context.NONE);
    }
}
