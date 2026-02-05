
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
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresCreateWithLocalAuthDisabled.json
     */
    /**
     * Sample code: ConfigurationStores_Create_With_Local_Auth_Disabled.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresCreateWithLocalAuthDisabled(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().define("contoso").withRegion("westus")
            .withExistingResourceGroup("myResourceGroup").withSku(new Sku().withName("Standard"))
            .withDisableLocalAuth(true)
            .withDataPlaneProxy(new DataPlaneProxyProperties().withAuthenticationMode(AuthenticationMode.PASS_THROUGH)
                .withPrivateLinkDelegation(PrivateLinkDelegation.DISABLED))
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
