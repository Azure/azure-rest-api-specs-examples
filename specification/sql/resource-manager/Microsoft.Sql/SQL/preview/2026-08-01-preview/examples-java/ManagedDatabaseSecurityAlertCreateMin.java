
import com.azure.resourcemanager.sql.fluent.models.ManagedDatabaseSecurityAlertPolicyInner;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyName;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyState;

/**
 * Samples for ManagedDatabaseSecurityAlertPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ManagedDatabaseSecurityAlertCreateMin.json
     */
    /**
     * Sample code: Update a database's threat detection policy with minimal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateADatabaseSThreatDetectionPolicyWithMinimalParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseSecurityAlertPolicies().createOrUpdateWithResponse(
            "securityalert-4799", "securityalert-6440", "testdb", SecurityAlertPolicyName.DEFAULT,
            new ManagedDatabaseSecurityAlertPolicyInner().withState(SecurityAlertPolicyState.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
