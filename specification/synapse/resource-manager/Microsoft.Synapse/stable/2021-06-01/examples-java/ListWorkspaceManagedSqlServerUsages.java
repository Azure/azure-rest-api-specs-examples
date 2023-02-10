/** Samples for WorkspaceManagedSqlServerUsages List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListWorkspaceManagedSqlServerUsages.json
     */
    /**
     * Sample code: List usages metric for the workspace managed sql server.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listUsagesMetricForTheWorkspaceManagedSqlServer(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.workspaceManagedSqlServerUsages().list("wsg-7398", "testWorkspace", com.azure.core.util.Context.NONE);
    }
}
