
import com.azure.resourcemanager.appconfiguration.models.AzureFrontDoorProperties;
import com.azure.resourcemanager.appconfiguration.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ConfigurationStores Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresCreateWithAzureFrontDoor.json
     */
    /**
     * Sample code: ConfigurationStores_Create_With_AzureFrontDoor.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresCreateWithAzureFrontDoor(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().define("contoso").withRegion("westus")
            .withExistingResourceGroup("myResourceGroup").withSku(new Sku().withName("Standard"))
            .withTags(mapOf("myTag", "myTagValue"))
            .withAzureFrontDoor(new AzureFrontDoorProperties().withResourceId(
                "/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/microsoft.cdn/profiles/myAzureFrontDoorProfile"))
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
