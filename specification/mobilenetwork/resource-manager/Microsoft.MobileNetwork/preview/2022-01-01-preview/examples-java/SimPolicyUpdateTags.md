```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.mobilenetwork.models.SimPolicy;
import java.util.HashMap;
import java.util.Map;

/** Samples for SimPolicies UpdateTags. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/SimPolicyUpdateTags.json
     */
    /**
     * Sample code: Update sim policy tags.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void updateSimPolicyTags(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        SimPolicy resource =
            manager.simPolicies().getWithResponse("rg1", "testMobileNetwork", "testPolicy", Context.NONE).getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).apply();
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mobilenetwork_1.0.0-beta.1/sdk/mobilenetwork/azure-resourcemanager-mobilenetwork/README.md) on how to add the SDK to your project and authenticate.
