import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.fluent.models.DedicatedSQLminimalTlsSettingsInner;
import com.azure.resourcemanager.synapse.models.DedicatedSqlMinimalTlsSettingsName;

/** Samples for WorkspaceManagedSqlServerDedicatedSqlMinimalTlsSettings Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/UpdateWorkspaceManagedSqlServerDedicatedSQLminimalTlsSettings.json
     */
    /**
     * Sample code: Update tls version of the workspace managed sql server.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void updateTlsVersionOfTheWorkspaceManagedSqlServer(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceManagedSqlServerDedicatedSqlMinimalTlsSettings()
            .update(
                "workspace-6852",
                "workspace-2080",
                DedicatedSqlMinimalTlsSettingsName.DEFAULT,
                new DedicatedSQLminimalTlsSettingsInner().withMinimalTlsVersion("1.1"),
                Context.NONE);
    }
}
