import com.azure.core.util.Context;
import com.azure.resourcemanager.iothub.models.RouteProperties;
import com.azure.resourcemanager.iothub.models.RoutingMessage;
import com.azure.resourcemanager.iothub.models.RoutingSource;
import com.azure.resourcemanager.iothub.models.TestRouteInput;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for IotHubResource TestRoute. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2022-04-30-preview/examples/iothub_testnewroute.json
     */
    /**
     * Sample code: IotHubResource_TestRoute.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceTestRoute(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .iotHubResources()
            .testRouteWithResponse(
                "testHub",
                "myResourceGroup",
                new TestRouteInput()
                    .withMessage(
                        new RoutingMessage()
                            .withBody("Body of message")
                            .withAppProperties(mapOf("key1", "value1"))
                            .withSystemProperties(mapOf("key1", "value1")))
                    .withRoute(
                        new RouteProperties()
                            .withName("Routeid")
                            .withSource(RoutingSource.DEVICE_MESSAGES)
                            .withEndpointNames(Arrays.asList("id1"))
                            .withIsEnabled(true)),
                Context.NONE);
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
