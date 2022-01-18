Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.fluent.models.StringDictionaryInner;
import java.util.HashMap;
import java.util.Map;

/** Samples for StaticSites CreateOrUpdateStaticSiteBuildFunctionAppSettings. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/CreateOrUpdateStaticSiteBuildFunctionAppSettings.json
     */
    /**
     * Sample code: Creates or updates the function app settings of a static site build.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsOrUpdatesTheFunctionAppSettingsOfAStaticSiteBuild(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .createOrUpdateStaticSiteBuildFunctionAppSettingsWithResponse(
                "rg",
                "testStaticSite0",
                "12",
                new StringDictionaryInner().withProperties(mapOf("setting1", "someval", "setting2", "someval2")),
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
