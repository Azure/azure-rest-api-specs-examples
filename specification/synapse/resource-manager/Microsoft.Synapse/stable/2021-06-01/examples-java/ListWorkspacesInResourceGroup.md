Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Workspaces ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListWorkspacesInResourceGroup.json
     */
    /**
     * Sample code: List workspaces in resource group.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listWorkspacesInResourceGroup(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.workspaces().listByResourceGroup("resourceGroup1", Context.NONE);
    }
}
```
