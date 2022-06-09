```java
import com.azure.core.util.Context;

/** Samples for WorkspaceManagedSqlServerDedicatedSqlMinimalTlsSettings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceManagedSqlServerDedicatedSQLminimalTlsSettings.json
     */
    /**
     * Sample code: Get workspace managed sql server dedicated sql minimal tls settings.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getWorkspaceManagedSqlServerDedicatedSqlMinimalTlsSettings(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceManagedSqlServerDedicatedSqlMinimalTlsSettings()
            .getWithResponse("workspace-6852", "workspace-2080", "default", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
