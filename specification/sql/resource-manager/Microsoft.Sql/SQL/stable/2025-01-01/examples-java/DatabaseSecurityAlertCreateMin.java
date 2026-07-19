
import com.azure.resourcemanager.sql.fluent.models.DatabaseSecurityAlertPolicyInner;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyName;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyState;

/**
 * Samples for DatabaseSecurityAlertPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseSecurityAlertCreateMin.json
     */
    /**
     * Sample code: Update a database's threat detection policy with minimal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateADatabaseSThreatDetectionPolicyWithMinimalParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseSecurityAlertPolicies().createOrUpdateWithResponse("securityalert-4799",
            "securityalert-6440", "testdb", SecurityAlertPolicyName.DEFAULT,
            new DatabaseSecurityAlertPolicyInner().withState(SecurityAlertPolicyState.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
