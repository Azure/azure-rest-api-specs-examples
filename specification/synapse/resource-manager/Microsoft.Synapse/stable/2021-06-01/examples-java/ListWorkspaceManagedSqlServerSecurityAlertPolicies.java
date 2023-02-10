/** Samples for WorkspaceManagedSqlServerSecurityAlertPolicy List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListWorkspaceManagedSqlServerSecurityAlertPolicies.json
     */
    /**
     * Sample code: Get workspace managed sql server's security alert policy.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getWorkspaceManagedSqlServerSSecurityAlertPolicy(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceManagedSqlServerSecurityAlertPolicies()
            .list("wsg-7398", "testWorkspace", com.azure.core.util.Context.NONE);
    }
}
