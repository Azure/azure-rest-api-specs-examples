
import com.azure.resourcemanager.appconfiguration.models.ConfigurationStore;
import com.azure.resourcemanager.appconfiguration.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ConfigurationStores Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresUpdate.json
     */
    /**
     * Sample code: ConfigurationStores_Update.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void
        configurationStoresUpdate(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        ConfigurationStore resource = manager.configurationStores()
            .getByResourceGroupWithResponse("myResourceGroup", "contoso", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("Category", "Marketing")).withSku(new Sku().withName("Standard")).apply();
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
