```java
import com.azure.core.util.Context;

/** Samples for BigDataPools Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/GetBigDataPool.json
     */
    /**
     * Sample code: Get a Big Data pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getABigDataPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.bigDataPools().getWithResponse("ExampleResourceGroup", "ExampleWorkspace", "ExamplePool", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
