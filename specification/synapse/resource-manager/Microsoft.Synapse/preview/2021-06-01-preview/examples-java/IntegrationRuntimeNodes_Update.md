```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.UpdateIntegrationRuntimeNodeRequest;

/** Samples for IntegrationRuntimeNodes Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimeNodes_Update.json
     */
    /**
     * Sample code: Update integration runtime node.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void updateIntegrationRuntimeNode(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .integrationRuntimeNodes()
            .updateWithResponse(
                "exampleResourceGroup",
                "exampleWorkspace",
                "exampleIntegrationRuntime",
                "Node_1",
                new UpdateIntegrationRuntimeNodeRequest().withConcurrentJobsLimit(2),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
