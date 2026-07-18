
import com.azure.resourcemanager.sql.fluent.models.ServerSecurityAlertPolicyInner;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyName;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyState;

/**
 * Samples for ServerSecurityAlertPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerSecurityAlertsCreateMin.json
     */
    /**
     * Sample code: Update a server's threat detection policy with minimal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateAServerSThreatDetectionPolicyWithMinimalParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerSecurityAlertPolicies().createOrUpdate("securityalert-4799",
            "securityalert-6440", SecurityAlertPolicyName.DEFAULT,
            new ServerSecurityAlertPolicyInner().withState(SecurityAlertPolicyState.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
