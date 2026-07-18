
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyName;

/**
 * Samples for ServerSecurityAlertPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerSecurityAlertsGet.json
     */
    /**
     * Sample code: Get a server's threat detection policy.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAServerSThreatDetectionPolicy(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerSecurityAlertPolicies().getWithResponse("securityalert-4799",
            "securityalert-6440", SecurityAlertPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
