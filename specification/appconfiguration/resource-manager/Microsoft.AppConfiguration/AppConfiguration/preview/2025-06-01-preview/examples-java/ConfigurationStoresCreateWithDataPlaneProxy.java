
import com.azure.resourcemanager.appconfiguration.models.AuthenticationMode;
import com.azure.resourcemanager.appconfiguration.models.DataPlaneProxyProperties;
import com.azure.resourcemanager.appconfiguration.models.PrivateLinkDelegation;
import com.azure.resourcemanager.appconfiguration.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ConfigurationStores Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresCreateWithDataPlaneProxy.json
     */
    /**
     * Sample code: ConfigurationStores_Create_With_Data_Plane_Proxy.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresCreateWithDataPlaneProxy(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().define("contoso").withRegion("westus")
            .withExistingResourceGroup("myResourceGroup").withSku(new Sku().withName("Standard"))
            .withDataPlaneProxy(new DataPlaneProxyProperties().withAuthenticationMode(AuthenticationMode.PASS_THROUGH)
                .withPrivateLinkDelegation(PrivateLinkDelegation.ENABLED))
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
