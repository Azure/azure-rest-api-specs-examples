
/**
 * Samples for SqlPoolSecurityAlertPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * SqlPoolSecurityAlertPolicies_List.json
     */
    /**
     * Sample code: Get a security alert of a SQL Analytics pool.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void getASecurityAlertOfASQLAnalyticsPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPoolSecurityAlertPolicies().list("securityalert-6852", "securityalert-2080", "testdb",
            com.azure.core.util.Context.NONE);
    }
}
