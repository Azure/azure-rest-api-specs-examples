
import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceSkuProperties;
import com.azure.resourcemanager.apimanagement.models.ConfigurationApi;
import com.azure.resourcemanager.apimanagement.models.LegacyApiState;
import com.azure.resourcemanager.apimanagement.models.SkuType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ApiManagementService CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateServiceWithoutLegacyConfigurationApi.json
     */
    /**
     * Sample code: ApiManagementCreateServiceWithoutLegacyConfigurationApi.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateServiceWithoutLegacyConfigurationApi(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiManagementServices().define("apimService1").withRegion("Central US").withExistingResourceGroup("rg1")
            .withSku(new ApiManagementServiceSkuProperties().withName(SkuType.BASIC).withCapacity(1))
            .withPublisherEmail("apim@autorestsdk.com").withPublisherName("autorestsdk")
            .withTags(mapOf("tag1", "value1", "tag2", "value2", "tag3", "value3"))
            .withConfigurationApi(new ConfigurationApi().withLegacyApi(LegacyApiState.DISABLED)).create();
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
