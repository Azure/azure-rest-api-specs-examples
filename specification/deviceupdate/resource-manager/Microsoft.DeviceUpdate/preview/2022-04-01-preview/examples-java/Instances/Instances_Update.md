```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.deviceupdate.models.Instance;
import java.util.HashMap;
import java.util.Map;

/** Samples for Instances Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/Instances/Instances_Update.json
     */
    /**
     * Sample code: Updates Instance.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void updatesInstance(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        Instance resource = manager.instances().getWithResponse("test-rg", "contoso", "blue", Context.NONE).getValue();
        resource.update().withTags(mapOf("tagKey", "tagValue")).apply();
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-deviceupdate_1.0.0-beta.1/sdk/deviceupdate/azure-resourcemanager-deviceupdate/README.md) on how to add the SDK to your project and authenticate.
