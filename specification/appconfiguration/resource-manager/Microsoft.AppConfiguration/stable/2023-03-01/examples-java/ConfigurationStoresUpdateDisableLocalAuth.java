import com.azure.resourcemanager.appconfiguration.models.ConfigurationStore;
import com.azure.resourcemanager.appconfiguration.models.Sku;
import java.util.HashMap;
import java.util.Map;

/** Samples for ConfigurationStores Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/ConfigurationStoresUpdateDisableLocalAuth.json
     */
    /**
     * Sample code: ConfigurationStores_Update_Disable_Local_Auth.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresUpdateDisableLocalAuth(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        ConfigurationStore resource =
            manager
                .configurationStores()
                .getByResourceGroupWithResponse("myResourceGroup", "contoso", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withSku(new Sku().withName("Standard")).withDisableLocalAuth(true).apply();
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
