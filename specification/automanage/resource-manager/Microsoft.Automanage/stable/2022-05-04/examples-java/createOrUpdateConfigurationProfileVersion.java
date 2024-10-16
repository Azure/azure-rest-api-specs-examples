
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.automanage.fluent.models.ConfigurationProfileInner;
import com.azure.resourcemanager.automanage.models.ConfigurationProfileProperties;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ConfigurationProfilesVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/
     * createOrUpdateConfigurationProfileVersion.json
     */
    /**
     * Sample code: Create or update configuration profile version.
     * 
     * @param manager Entry point to AutomanageManager.
     */
    public static void createOrUpdateConfigurationProfileVersion(
        com.azure.resourcemanager.automanage.AutomanageManager manager) throws IOException {
        manager.configurationProfilesVersions().createOrUpdateWithResponse("customConfigurationProfile", "version1",
            "myResourceGroupName",
            new ConfigurationProfileInner().withLocation("East US").withTags(mapOf("Organization", "Administration"))
                .withProperties(new ConfigurationProfileProperties()
                    .withConfiguration(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                        "{\"Antimalware/Enable\":false,\"AzureSecurityCenter/Enable\":true,\"Backup/Enable\":false,\"BootDiagnostics/Enable\":true,\"ChangeTrackingAndInventory/Enable\":true,\"GuestConfiguration/Enable\":true,\"LogAnalytics/Enable\":true,\"UpdateManagement/Enable\":true,\"VMInsights/Enable\":true}",
                        Object.class, SerializerEncoding.JSON))),
            com.azure.core.util.Context.NONE);
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
