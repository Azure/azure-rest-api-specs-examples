
/**
 * Samples for WorkspaceManagedSqlServerDedicatedSqlMinimalTlsSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/
     * GetWorkspaceManagedSqlServerDedicatedSQLminimalTlsSettings.json
     */
    /**
     * Sample code: Get workspace managed sql server dedicated sql minimal tls settings.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void getWorkspaceManagedSqlServerDedicatedSqlMinimalTlsSettings(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.workspaceManagedSqlServerDedicatedSqlMinimalTlsSettings().getWithResponse("workspace-6852",
            "workspace-2080", "default", com.azure.core.util.Context.NONE);
    }
}
