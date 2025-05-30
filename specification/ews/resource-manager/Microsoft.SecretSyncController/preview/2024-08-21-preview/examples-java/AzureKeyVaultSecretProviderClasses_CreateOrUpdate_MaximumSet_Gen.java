
import com.azure.resourcemanager.secretsstoreextension.models.AzureKeyVaultSecretProviderClassProperties;
import com.azure.resourcemanager.secretsstoreextension.models.ExtendedLocation;
import com.azure.resourcemanager.secretsstoreextension.models.ExtendedLocationType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AzureKeyVaultSecretProviderClasses CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-21-preview/AzureKeyVaultSecretProviderClasses_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: AzureKeyVaultSecretProviderClasses_CreateOrUpdate.
     * 
     * @param manager Entry point to SecretsStoreExtensionManager.
     */
    public static void azureKeyVaultSecretProviderClassesCreateOrUpdate(
        com.azure.resourcemanager.secretsstoreextension.SecretsStoreExtensionManager manager) {
        manager.azureKeyVaultSecretProviderClasses().define("akvspc-ssc-example").withRegion("eastus")
            .withExistingResourceGroup("rg-ssc-example").withTags(mapOf("example-tag", "example-tag-value"))
            .withProperties(new AzureKeyVaultSecretProviderClassProperties().withKeyvaultName("fakeTokenPlaceholder")
                .withClientId("00000000-0000-0000-0000-000000000000")
                .withTenantId("00000000-0000-0000-0000-000000000000").withObjects(
                    "array: |\n  - |\n    objectName: my-secret-object\n    objectType: secret\n    objectVersionHistory: 1"))
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-ssc-example/providers/Microsoft.ExtendedLocation/customLocations/example-custom-location")
                .withType(ExtendedLocationType.CUSTOM_LOCATION))
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
