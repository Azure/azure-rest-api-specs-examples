
/**
 * Samples for ManagedServerSecurityAlertPolicies ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedServerSecurityAlertListByInstance.json
     */
    /**
     * Sample code: Get the managed server's threat detection policies.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getTheManagedServerSThreatDetectionPolicies(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedServerSecurityAlertPolicies().listByInstance("securityalert-4799",
            "securityalert-6440", com.azure.core.util.Context.NONE);
    }
}
