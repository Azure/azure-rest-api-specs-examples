
import com.azure.core.util.BinaryData;
import com.azure.resourcemanager.deviceregistry.models.DeviceMessagingEndpoint;
import com.azure.resourcemanager.deviceregistry.models.MessagingEndpoints;
import com.azure.resourcemanager.deviceregistry.models.NamespaceDeviceProperties;
import com.azure.resourcemanager.deviceregistry.models.OutboundEndpoints;
import java.nio.charset.StandardCharsets;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NamespaceDevices CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/CreateOrReplace_NamespaceDevice.json
     */
    /**
     * Sample code: CreateOrReplace_NamespaceDevices.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        createOrReplaceNamespaceDevices(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaceDevices().define("dev-namespace-gbk0925-n01").withRegion("West Europe")
            .withExistingNamespace("myResourceGroup", "adr-namespace-gbk0925-n01")
            .withProperties(new NamespaceDeviceProperties().withEnabled(true)
                .withExternalDeviceId("adr-smart-device3-7a848b15-af47-40a7-8c06-a3f43314d44f")
                .withEndpoints(new MessagingEndpoints()
                    .withOutbound(new OutboundEndpoints().withAssigned(mapOf("eventGridEndpoint",
                        new DeviceMessagingEndpoint().withEndpointType("Microsoft.Devices/IoTHubs")
                            .withAddress("https://myeventgridtopic.westeurope-1.eventgrid.azure.net/api/events")))))
                .withAttributes(mapOf("deviceType", BinaryData.fromBytes("sensor".getBytes(StandardCharsets.UTF_8)),
                    "deviceOwner", BinaryData.fromBytes("IT".getBytes(StandardCharsets.UTF_8)), "deviceCategory",
                    BinaryData.fromBytes("16".getBytes(StandardCharsets.UTF_8)))))
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
