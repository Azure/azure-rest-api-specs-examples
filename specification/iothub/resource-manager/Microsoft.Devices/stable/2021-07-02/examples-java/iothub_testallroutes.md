Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-iothub_1.2.0-beta.1/sdk/iothub/azure-resourcemanager-iothub/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.iothub.models.RoutingMessage;
import com.azure.resourcemanager.iothub.models.RoutingSource;
import com.azure.resourcemanager.iothub.models.TestAllRoutesInput;
import java.util.HashMap;
import java.util.Map;

/** Samples for IotHubResource TestAllRoutes. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_testallroutes.json
     */
    /**
     * Sample code: IotHubResource_TestAllRoutes.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceTestAllRoutes(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .iotHubResources()
            .testAllRoutesWithResponse(
                "testHub",
                "myResourceGroup",
                new TestAllRoutesInput()
                    .withRoutingSource(RoutingSource.DEVICE_MESSAGES)
                    .withMessage(
                        new RoutingMessage()
                            .withBody("Body of message")
                            .withAppProperties(mapOf("key1", "value1"))
                            .withSystemProperties(mapOf("key1", "value1"))),
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
```
