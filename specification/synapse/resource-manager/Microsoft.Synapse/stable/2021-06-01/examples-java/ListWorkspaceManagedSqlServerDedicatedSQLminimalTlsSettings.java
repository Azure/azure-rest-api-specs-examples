/** Samples for WorkspaceManagedSqlServerDedicatedSqlMinimalTlsSettings List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListWorkspaceManagedSqlServerDedicatedSQLminimalTlsSettings.json
     */
    /**
     * Sample code: List dedicated sql minimal tls settings of the workspace managed sql server.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listDedicatedSqlMinimalTlsSettingsOfTheWorkspaceManagedSqlServer(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceManagedSqlServerDedicatedSqlMinimalTlsSettings()
            .list("workspace-6852", "workspace-2080", com.azure.core.util.Context.NONE);
    }
}
