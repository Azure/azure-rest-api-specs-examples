
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.automanage.models.ConfigurationProfile;
import com.azure.resourcemanager.automanage.models.ConfigurationProfileProperties;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ConfigurationProfiles Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/
     * updateConfigurationProfile.json
     */
    /**
     * Sample code: Update configuration profile.
     * 
     * @param manager Entry point to AutomanageManager.
     */
    public static void updateConfigurationProfile(com.azure.resourcemanager.automanage.AutomanageManager manager)
        throws IOException {
        ConfigurationProfile resource
            = manager.configurationProfiles().getByResourceGroupWithResponse("myResourceGroupName",
                "customConfigurationProfile", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("Organization", "Administration"))
            .withProperties(new ConfigurationProfileProperties()
                .withConfiguration(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                    "{\"Antimalware/Enable\":false,\"AzureSecurityCenter/Enable\":true,\"Backup/Enable\":false,\"BootDiagnostics/Enable\":true,\"ChangeTrackingAndInventory/Enable\":true,\"GuestConfiguration/Enable\":true,\"LogAnalytics/Enable\":true,\"UpdateManagement/Enable\":true,\"VMInsights/Enable\":true}",
                    Object.class, SerializerEncoding.JSON)))
            .apply();
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
