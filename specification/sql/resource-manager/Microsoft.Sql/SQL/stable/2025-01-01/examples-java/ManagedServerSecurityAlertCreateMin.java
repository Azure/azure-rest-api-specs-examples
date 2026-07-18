
import com.azure.resourcemanager.sql.fluent.models.ManagedServerSecurityAlertPolicyInner;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyName;
import com.azure.resourcemanager.sql.models.SecurityAlertPolicyState;

/**
 * Samples for ManagedServerSecurityAlertPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedServerSecurityAlertCreateMin.json
     */
    /**
     * Sample code: Update a managed server's threat detection policy with minimal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateAManagedServerSThreatDetectionPolicyWithMinimalParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedServerSecurityAlertPolicies().createOrUpdate("securityalert-4799",
            "securityalert-6440", SecurityAlertPolicyName.DEFAULT,
            new ManagedServerSecurityAlertPolicyInner().withState(SecurityAlertPolicyState.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
