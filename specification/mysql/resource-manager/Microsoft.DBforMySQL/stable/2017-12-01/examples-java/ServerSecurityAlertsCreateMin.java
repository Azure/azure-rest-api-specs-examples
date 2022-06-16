import com.azure.core.util.Context;
import com.azure.resourcemanager.mysql.models.SecurityAlertPolicyName;
import com.azure.resourcemanager.mysql.models.ServerSecurityAlertPolicy;
import com.azure.resourcemanager.mysql.models.ServerSecurityAlertPolicyState;

/** Samples for ServerSecurityAlertPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerSecurityAlertsCreateMin.json
     */
    /**
     * Sample code: Update a server's threat detection policy with minimal parameters.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void updateAServerSThreatDetectionPolicyWithMinimalParameters(
        com.azure.resourcemanager.mysql.MySqlManager manager) {
        ServerSecurityAlertPolicy resource =
            manager
                .serverSecurityAlertPolicies()
                .getWithResponse(
                    "securityalert-4799", "securityalert-6440", SecurityAlertPolicyName.DEFAULT, Context.NONE)
                .getValue();
        resource.update().withState(ServerSecurityAlertPolicyState.DISABLED).withEmailAccountAdmins(true).apply();
    }
}
