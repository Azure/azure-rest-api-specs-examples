
import com.azure.resourcemanager.appconfiguration.models.IdentityType;
import com.azure.resourcemanager.appconfiguration.models.ResourceIdentity;
import com.azure.resourcemanager.appconfiguration.models.Sku;
import com.azure.resourcemanager.appconfiguration.models.UserIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ConfigurationStores Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresCreateWithIdentity.json
     */
    /**
     * Sample code: ConfigurationStores_Create_With_Identity.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresCreateWithIdentity(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().define("contoso").withRegion("westus")
            .withExistingResourceGroup("myResourceGroup").withSku(new Sku().withName("Standard"))
            .withTags(mapOf("myTag", "myTagValue"))
            .withIdentity(new ResourceIdentity().withType(IdentityType.SYSTEM_ASSIGNED_USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourcegroups/myResourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2",
                    new UserIdentity())))
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
