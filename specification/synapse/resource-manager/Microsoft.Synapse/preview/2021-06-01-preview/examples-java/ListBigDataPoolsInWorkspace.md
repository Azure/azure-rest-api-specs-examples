```java
import com.azure.core.util.Context;

/** Samples for BigDataPools ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/ListBigDataPoolsInWorkspace.json
     */
    /**
     * Sample code: List Big Data pools in a workspace.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listBigDataPoolsInAWorkspace(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.bigDataPools().listByWorkspace("ExampleResourceGroup", "ExampleWorkspace", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
