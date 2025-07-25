
import com.azure.resourcemanager.appconfiguration.models.ConfigurationStore;
import com.azure.resourcemanager.appconfiguration.models.IdentityType;
import com.azure.resourcemanager.appconfiguration.models.ResourceIdentity;
import com.azure.resourcemanager.appconfiguration.models.Sku;
import com.azure.resourcemanager.appconfiguration.models.UserIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ConfigurationStores Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2024-06-01/examples/
     * ConfigurationStoresUpdateWithIdentity.json
     */
    /**
     * Sample code: ConfigurationStores_Update_With_Identity.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresUpdateWithIdentity(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        ConfigurationStore resource = manager.configurationStores()
            .getByResourceGroupWithResponse("myResourceGroup", "contoso", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("Category", "Marketing")).withIdentity(new ResourceIdentity()
            .withType(IdentityType.SYSTEM_ASSIGNED_USER_ASSIGNED)
            .withUserAssignedIdentities(mapOf(
                "/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourcegroups/myResourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2",
                new UserIdentity())))
            .withSku(new Sku().withName("Standard")).apply();
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
