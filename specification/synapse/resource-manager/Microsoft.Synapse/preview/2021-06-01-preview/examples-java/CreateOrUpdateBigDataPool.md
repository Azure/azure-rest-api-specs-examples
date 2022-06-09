```java
import com.azure.resourcemanager.synapse.models.AutoPauseProperties;
import com.azure.resourcemanager.synapse.models.AutoScaleProperties;
import com.azure.resourcemanager.synapse.models.LibraryRequirements;
import com.azure.resourcemanager.synapse.models.NodeSize;
import com.azure.resourcemanager.synapse.models.NodeSizeFamily;
import java.util.HashMap;
import java.util.Map;

/** Samples for BigDataPools CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/CreateOrUpdateBigDataPool.json
     */
    /**
     * Sample code: Create or update a Big Data pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateABigDataPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .bigDataPools()
            .define("ExamplePool")
            .withRegion("West US 2")
            .withExistingWorkspace("ExampleResourceGroup", "ExampleWorkspace")
            .withTags(mapOf("key", "value"))
            .withAutoScale(new AutoScaleProperties().withMinNodeCount(3).withEnabled(true).withMaxNodeCount(50))
            .withAutoPause(new AutoPauseProperties().withDelayInMinutes(15).withEnabled(true))
            .withSparkEventsFolder("/events")
            .withNodeCount(4)
            .withLibraryRequirements(new LibraryRequirements().withContent("").withFilename("requirements.txt"))
            .withSparkVersion("2.4")
            .withDefaultSparkLogFolder("/logs")
            .withNodeSize(NodeSize.MEDIUM)
            .withNodeSizeFamily(NodeSizeFamily.MEMORY_OPTIMIZED)
            .create();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
