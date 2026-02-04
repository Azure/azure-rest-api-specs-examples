
import com.azure.resourcemanager.appconfiguration.models.Sku;
import com.azure.resourcemanager.appconfiguration.models.TelemetryProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ConfigurationStores Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresCreateWithTelemetry.json
     */
    /**
     * Sample code: ConfigurationStores_Create_With_Telemetry.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresCreateWithTelemetry(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().define("contoso").withRegion("westus")
            .withExistingResourceGroup("myResourceGroup").withSku(new Sku().withName("Standard"))
            .withTags(mapOf("myTag", "myTagValue"))
            .withTelemetry(new TelemetryProperties().withResourceId(
                "/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/microsoft.insights/components/appInsightsName"))
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
