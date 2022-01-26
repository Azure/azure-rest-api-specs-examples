Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for LibrariesOperation ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/Libraries_ListByWorkspace.json
     */
    /**
     * Sample code: List libraries in a workspace.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listLibrariesInAWorkspace(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.librariesOperations().listByWorkspace("exampleResourceGroup", "exampleWorkspace", Context.NONE);
    }
}
```
