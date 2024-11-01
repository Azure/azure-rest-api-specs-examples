
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.streamanalytics.models.CompatibilityLevel;
import com.azure.resourcemanager.streamanalytics.models.EventsOutOfOrderPolicy;
import com.azure.resourcemanager.streamanalytics.models.Identity;
import com.azure.resourcemanager.streamanalytics.models.OutputErrorPolicy;
import com.azure.resourcemanager.streamanalytics.models.Sku;
import com.azure.resourcemanager.streamanalytics.models.SkuName;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StreamingJobs CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * StreamingJob_Create_UserAssignedIdentity.json
     */
    /**
     * Sample code: Create a streaming job shell (a streaming job with no inputs, outputs, transformation, or functions)
     * with user assigned identity.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        createAStreamingJobShellAStreamingJobWithNoInputsOutputsTransformationOrFunctionsWithUserAssignedIdentity(
            com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) throws IOException {
        manager.streamingJobs().define("sjName").withRegion("West US").withExistingResourceGroup("sjrg")
            .withTags(mapOf("key1", "fakeTokenPlaceholder", "key3", "fakeTokenPlaceholder", "randomKey",
                "fakeTokenPlaceholder"))
            .withIdentity(new Identity().withType("UserAssigned").withUserAssignedIdentities(mapOf(
                "/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.ManagedIdentity/userAssignedIdentities/sdkIdentity",
                SerializerFactory.createDefaultManagementSerializerAdapter().deserialize("{}", Object.class,
                    SerializerEncoding.JSON))))
            .withSkuPropertiesSku(new Sku().withName(SkuName.STANDARD))
            .withEventsOutOfOrderPolicy(EventsOutOfOrderPolicy.DROP).withOutputErrorPolicy(OutputErrorPolicy.DROP)
            .withEventsOutOfOrderMaxDelayInSeconds(5).withEventsLateArrivalMaxDelayInSeconds(16).withDataLocale("en-US")
            .withCompatibilityLevel(CompatibilityLevel.ONE_ZERO).withInputs(Arrays.asList())
            .withOutputs(Arrays.asList()).withFunctions(Arrays.asList()).create();
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
