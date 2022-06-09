```java
import com.azure.core.util.Context;

/** Samples for WorkspaceManagedSqlServerBlobAuditingPolicies ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListWorkspaceManagedSqlServerBlobAuditingSettings.json
     */
    /**
     * Sample code: Get blob auditing policy of workspace manged sql Server.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getBlobAuditingPolicyOfWorkspaceMangedSqlServer(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceManagedSqlServerBlobAuditingPolicies()
            .listByWorkspace("wsg-7398", "testWorkspace", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
