
import com.azure.resourcemanager.postgresql.models.SecurityAlertPolicyName;
import com.azure.resourcemanager.postgresql.models.ServerSecurityAlertPolicy;
import com.azure.resourcemanager.postgresql.models.ServerSecurityAlertPolicyState;

/**
 * Samples for ServerSecurityAlertPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/
     * ServerSecurityAlertsCreateMin.json
     */
    /**
     * Sample code: Update a server's threat detection policy with minimal parameters.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void updateAServerSThreatDetectionPolicyWithMinimalParameters(
        com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        ServerSecurityAlertPolicy resource = manager.serverSecurityAlertPolicies().getWithResponse("securityalert-4799",
            "securityalert-6440", SecurityAlertPolicyName.DEFAULT, com.azure.core.util.Context.NONE).getValue();
        resource.update().withState(ServerSecurityAlertPolicyState.DISABLED).withEmailAccountAdmins(true).apply();
    }
}
