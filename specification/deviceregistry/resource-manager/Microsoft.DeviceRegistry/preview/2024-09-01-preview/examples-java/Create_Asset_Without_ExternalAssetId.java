
import com.azure.resourcemanager.deviceregistry.models.AssetProperties;
import com.azure.resourcemanager.deviceregistry.models.DataPoint;
import com.azure.resourcemanager.deviceregistry.models.DataPointObservabilityMode;
import com.azure.resourcemanager.deviceregistry.models.Dataset;
import com.azure.resourcemanager.deviceregistry.models.Event;
import com.azure.resourcemanager.deviceregistry.models.EventObservabilityMode;
import com.azure.resourcemanager.deviceregistry.models.ExtendedLocation;
import com.azure.resourcemanager.deviceregistry.models.Topic;
import com.azure.resourcemanager.deviceregistry.models.TopicRetainType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Assets CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Create_Asset_Without_ExternalAssetId.json
     */
    /**
     * Sample code: Create_Asset_Without_ExternalAssetId.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        createAssetWithoutExternalAssetId(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.assets().define("my-asset").withRegion("West Europe").withExistingResourceGroup("myResourceGroup")
            .withExtendedLocation(new ExtendedLocation().withType("CustomLocation").withName(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"))
            .withTags(mapOf("site", "building-1"))
            .withProperties(new AssetProperties().withEnabled(true).withDisplayName("AssetDisplayName")
                .withDescription("This is a sample Asset").withAssetEndpointProfileRef("myAssetEndpointProfile")
                .withManufacturer("Contoso").withManufacturerUri("https://www.contoso.com/manufacturerUri")
                .withModel("ContosoModel").withProductCode("fakeTokenPlaceholder").withHardwareRevision("1.0")
                .withSoftwareRevision("2.0").withDocumentationUri("https://www.example.com/manual")
                .withSerialNumber("64-103816-519918-8")
                .withDefaultDatasetsConfiguration(
                    "{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}")
                .withDefaultEventsConfiguration("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}")
                .withDefaultTopic(new Topic().withPath("/path/defaultTopic").withRetain(TopicRetainType.KEEP))
                .withDatasets(Arrays.asList(new Dataset().withName("dataset1")
                    .withDatasetConfiguration("{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}")
                    .withTopic(new Topic().withPath("/path/dataset1").withRetain(TopicRetainType.KEEP))
                    .withDataPoints(Arrays.asList(
                        new DataPoint().withName("dataPoint1")
                            .withDataSource("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt1")
                            .withDataPointConfiguration(
                                "{\"publishingInterval\":8,\"samplingInterval\":8,\"queueSize\":4}")
                            .withObservabilityMode(DataPointObservabilityMode.COUNTER),
                        new DataPoint().withName("dataPoint2")
                            .withDataSource("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt2")
                            .withDataPointConfiguration(
                                "{\"publishingInterval\":4,\"samplingInterval\":4,\"queueSize\":7}")
                            .withObservabilityMode(DataPointObservabilityMode.NONE)))))
                .withEvents(
                    Arrays
                        .asList(
                            new Event().withName("event1")
                                .withEventNotifier("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt3")
                                .withEventConfiguration(
                                    "{\"publishingInterval\":7,\"samplingInterval\":1,\"queueSize\":8}")
                                .withTopic(new Topic().withPath("/path/event1").withRetain(TopicRetainType.KEEP))
                                .withObservabilityMode(EventObservabilityMode.NONE),
                            new Event().withName("event2")
                                .withEventNotifier("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt4")
                                .withEventConfiguration(
                                    "{\"publishingInterval\":7,\"samplingInterval\":8,\"queueSize\":4}")
                                .withObservabilityMode(EventObservabilityMode.LOG))))
            .create();
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
