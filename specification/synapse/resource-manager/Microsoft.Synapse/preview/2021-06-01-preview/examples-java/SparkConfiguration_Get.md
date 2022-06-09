```java
import com.azure.core.util.Context;

/** Samples for SparkConfiguration Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/SparkConfiguration_Get.json
     */
    /**
     * Sample code: Get SparkConfiguration by name.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getSparkConfigurationByName(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sparkConfigurations()
            .getWithResponse("exampleResourceGroup", "exampleSparkConfigurationName", "exampleWorkspace", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
