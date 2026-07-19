
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyName;

/**
 * Samples for ManagedDatabaseSecurityAlertPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseSecurityAlertGet.json
     */
    /**
     * Sample code: Get a database's threat detection policy.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getADatabaseSThreatDetectionPolicy(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseSecurityAlertPolicies().getWithResponse("securityalert-6852",
            "securityalert-2080", "testdb", SecurityAlertPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
