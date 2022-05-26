Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-deviceupdate_1.0.0-beta.1/sdk/deviceupdate/azure-resourcemanager-deviceupdate/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.deviceupdate.models.Account;
import java.util.HashMap;
import java.util.Map;

/** Samples for Accounts Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/Accounts/Accounts_Update.json
     */
    /**
     * Sample code: Updates Account.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void updatesAccount(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        Account resource =
            manager.accounts().getByResourceGroupWithResponse("test-rg", "contoso", Context.NONE).getValue();
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
