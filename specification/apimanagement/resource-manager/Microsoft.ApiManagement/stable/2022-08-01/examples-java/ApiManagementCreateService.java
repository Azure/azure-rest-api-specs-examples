import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceSkuProperties;
import com.azure.resourcemanager.apimanagement.models.SkuType;
import java.util.HashMap;
import java.util.Map;

/** Samples for ApiManagementService CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateService.json
     */
    /**
     * Sample code: ApiManagementCreateService.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateService(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiManagementServices()
            .define("apimService1")
            .withRegion("South Central US")
            .withExistingResourceGroup("rg1")
            .withSku(new ApiManagementServiceSkuProperties().withName(SkuType.DEVELOPER).withCapacity(1))
            .withPublisherEmail("foo@contoso.com")
            .withPublisherName("foo")
            .withTags(mapOf("Name", "Contoso", "Test", "User"))
            .create();
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
