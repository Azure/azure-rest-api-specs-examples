import com.azure.resourcemanager.mariadb.models.SecurityAlertPolicyName;
import com.azure.resourcemanager.mariadb.models.ServerSecurityAlertPolicy;
import com.azure.resourcemanager.mariadb.models.ServerSecurityAlertPolicyState;

/** Samples for ServerSecurityAlertPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerSecurityAlertsCreateMin.json
     */
    /**
     * Sample code: Update a server's threat detection policy with minimal parameters.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void updateAServerSThreatDetectionPolicyWithMinimalParameters(
        com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        ServerSecurityAlertPolicy resource =
            manager
                .serverSecurityAlertPolicies()
                .getWithResponse(
                    "securityalert-4799",
                    "securityalert-6440",
                    SecurityAlertPolicyName.DEFAULT,
                    com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withState(ServerSecurityAlertPolicyState.DISABLED).withEmailAccountAdmins(true).apply();
    }
}
