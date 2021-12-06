Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.fluent.models.WorkspaceAadAdminInfoInner;

/** Samples for WorkspaceSqlAadAdmins CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateWorkspaceAadAdmin.json
     */
    /**
     * Sample code: Create or update workspace active directory admin.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateWorkspaceActiveDirectoryAdmin(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceSqlAadAdmins()
            .createOrUpdate(
                "resourceGroup1",
                "workspace1",
                new WorkspaceAadAdminInfoInner()
                    .withTenantId("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c")
                    .withLogin("bob@contoso.com")
                    .withAdministratorType("ActiveDirectory")
                    .withSid("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
                Context.NONE);
    }
}
```
