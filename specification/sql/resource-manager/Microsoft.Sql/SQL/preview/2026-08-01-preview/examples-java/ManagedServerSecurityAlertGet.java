
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyName;

/**
 * Samples for ManagedServerSecurityAlertPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ManagedServerSecurityAlertGet.json
     */
    /**
     * Sample code: Get a managed server's threat detection policy.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAManagedServerSThreatDetectionPolicy(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedServerSecurityAlertPolicies().getWithResponse("securityalert-4799",
            "securityalert-6440", SecurityAlertPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
