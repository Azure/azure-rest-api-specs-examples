Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.streamanalytics.models.CompatibilityLevel;
import com.azure.resourcemanager.streamanalytics.models.EventsOutOfOrderPolicy;
import com.azure.resourcemanager.streamanalytics.models.OutputErrorPolicy;
import com.azure.resourcemanager.streamanalytics.models.Sku;
import com.azure.resourcemanager.streamanalytics.models.SkuName;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for StreamingJobs CreateOrReplace. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Create_JobShell.json
     */
    /**
     * Sample code: Create a streaming job shell (a streaming job with no inputs, outputs, transformation, or
     * functions).
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createAStreamingJobShellAStreamingJobWithNoInputsOutputsTransformationOrFunctions(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager
            .streamingJobs()
            .define("sj59")
            .withRegion("West US")
            .withExistingResourceGroup("sjrg6936")
            .withTags(mapOf("key1", "value1", "key3", "value3", "randomKey", "randomValue"))
            .withSku(new Sku().withName(SkuName.STANDARD))
            .withEventsOutOfOrderPolicy(EventsOutOfOrderPolicy.DROP)
            .withOutputErrorPolicy(OutputErrorPolicy.DROP)
            .withEventsOutOfOrderMaxDelayInSeconds(5)
            .withEventsLateArrivalMaxDelayInSeconds(16)
            .withDataLocale("en-US")
            .withCompatibilityLevel(CompatibilityLevel.ONE_ZERO)
            .withInputs(Arrays.asList())
            .withOutputs(Arrays.asList())
            .withFunctions(Arrays.asList())
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
