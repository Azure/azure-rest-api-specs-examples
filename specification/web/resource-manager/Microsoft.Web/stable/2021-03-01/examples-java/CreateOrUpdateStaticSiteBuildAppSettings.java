import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.fluent.models.StringDictionaryInner;
import java.util.HashMap;
import java.util.Map;

/** Samples for StaticSites CreateOrUpdateStaticSiteBuildAppSettings. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/CreateOrUpdateStaticSiteBuildAppSettings.json
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
            .createOrUpdateStaticSiteBuildAppSettingsWithResponse(
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
