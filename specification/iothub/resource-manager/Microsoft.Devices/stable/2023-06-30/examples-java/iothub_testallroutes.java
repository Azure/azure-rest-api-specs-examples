
import com.azure.resourcemanager.iothub.models.RoutingMessage;
import com.azure.resourcemanager.iothub.models.RoutingSource;
import com.azure.resourcemanager.iothub.models.TestAllRoutesInput;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for IotHubResource TestAllRoutes.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_testallroutes.json
     */
    /**
     * Sample code: IotHubResource_TestAllRoutes.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceTestAllRoutes(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().testAllRoutesWithResponse("testHub", "myResourceGroup",
            new TestAllRoutesInput().withRoutingSource(RoutingSource.DEVICE_MESSAGES)
                .withMessage(new RoutingMessage().withBody("Body of message")
                    .withAppProperties(mapOf("key1", "fakeTokenPlaceholder"))
                    .withSystemProperties(mapOf("key1", "fakeTokenPlaceholder"))),
            com.azure.core.util.Context.NONE);
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
