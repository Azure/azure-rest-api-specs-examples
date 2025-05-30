
import com.azure.resourcemanager.synapse.models.AutoPauseProperties;
import com.azure.resourcemanager.synapse.models.AutoScaleProperties;
import com.azure.resourcemanager.synapse.models.LibraryRequirements;
import com.azure.resourcemanager.synapse.models.NodeSize;
import com.azure.resourcemanager.synapse.models.NodeSizeFamily;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BigDataPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/CreateOrUpdateBigDataPool.
     * json
     */
    /**
     * Sample code: Create or update a Big Data pool.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateABigDataPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.bigDataPools().define("ExamplePool").withRegion("West US 2")
            .withExistingWorkspace("ExampleResourceGroup", "ExampleWorkspace")
            .withTags(mapOf("key", "fakeTokenPlaceholder"))
            .withAutoScale(new AutoScaleProperties().withMinNodeCount(3).withEnabled(true).withMaxNodeCount(50))
            .withAutoPause(new AutoPauseProperties().withDelayInMinutes(15).withEnabled(true))
            .withSparkEventsFolder("/events").withNodeCount(4)
            .withLibraryRequirements(new LibraryRequirements().withContent("").withFilename("requirements.txt"))
            .withSparkVersion("3.3").withDefaultSparkLogFolder("/logs").withNodeSize(NodeSize.MEDIUM)
            .withNodeSizeFamily(NodeSizeFamily.MEMORY_OPTIMIZED).create();
    }

    // Use "Map.of" if available
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
