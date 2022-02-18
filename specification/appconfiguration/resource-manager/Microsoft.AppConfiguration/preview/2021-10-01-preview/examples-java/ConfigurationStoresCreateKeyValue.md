Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-appconfiguration_1.0.0-beta.4/sdk/appconfiguration/azure-resourcemanager-appconfiguration/README.md) on how to add the SDK to your project and authenticate.

```java
import java.util.HashMap;
import java.util.Map;

/** Samples for KeyValues CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-10-01-preview/examples/ConfigurationStoresCreateKeyValue.json
     */
    /**
     * Sample code: KeyValues_CreateOrUpdate.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void keyValuesCreateOrUpdate(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager
            .keyValues()
            .define("myKey$myLabel")
            .withExistingConfigurationStore("myResourceGroup", "contoso")
            .withTags(mapOf("tag1", "tagValue1", "tag2", "tagValue2"))
            .withValue("myValue")
            .create();
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
