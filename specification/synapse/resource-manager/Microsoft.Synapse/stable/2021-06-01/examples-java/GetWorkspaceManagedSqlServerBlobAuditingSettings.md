Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.BlobAuditingPolicyName;

/** Samples for WorkspaceManagedSqlServerBlobAuditingPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceManagedSqlServerBlobAuditingSettings.json
     */
    /**
     * Sample code: Get blob auditing setting of workspace managed sql Server.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getBlobAuditingSettingOfWorkspaceManagedSqlServer(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceManagedSqlServerBlobAuditingPolicies()
            .getWithResponse("wsg-7398", "testWorkspace", BlobAuditingPolicyName.DEFAULT, Context.NONE);
    }
}
```
