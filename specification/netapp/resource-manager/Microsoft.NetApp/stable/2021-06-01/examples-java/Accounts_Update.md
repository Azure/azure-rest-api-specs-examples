Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-netapp_1.0.0-beta.6/sdk/netapp/azure-resourcemanager-netapp/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.netapp.models.NetAppAccount;
import java.util.HashMap;
import java.util.Map;

/** Samples for Accounts Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-06-01/examples/Accounts_Update.json
     */
    /**
     * Sample code: Accounts_Update.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountsUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        NetAppAccount resource =
            manager.accounts().getByResourceGroupWithResponse("myRG", "account1", Context.NONE).getValue();
        resource.update().withTags(mapOf("Tag1", "Value1")).apply();
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
