
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyName;

/**
 * Samples for DatabaseSecurityAlertPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseSecurityAlertGet.json
     */
    /**
     * Sample code: Get a database's threat detection policy.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getADatabaseSThreatDetectionPolicy(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseSecurityAlertPolicies().getWithResponse("securityalert-6852",
            "securityalert-2080", "testdb", SecurityAlertPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
