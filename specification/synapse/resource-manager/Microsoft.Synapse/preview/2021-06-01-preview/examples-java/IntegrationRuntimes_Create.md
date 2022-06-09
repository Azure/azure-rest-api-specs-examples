```java
import com.azure.resourcemanager.synapse.models.SelfHostedIntegrationRuntime;

/** Samples for IntegrationRuntimes Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Create.json
     */
    /**
     * Sample code: Create integration runtime.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createIntegrationRuntime(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .integrationRuntimes()
            .define("exampleIntegrationRuntime")
            .withExistingWorkspace("exampleResourceGroup", "exampleWorkspace")
            .withProperties(new SelfHostedIntegrationRuntime().withDescription("A selfhosted integration runtime"))
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
