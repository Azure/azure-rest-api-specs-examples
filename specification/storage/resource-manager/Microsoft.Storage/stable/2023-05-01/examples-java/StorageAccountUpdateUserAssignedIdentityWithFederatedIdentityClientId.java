
import com.azure.resourcemanager.storage.models.Encryption;
import com.azure.resourcemanager.storage.models.EncryptionIdentity;
import com.azure.resourcemanager.storage.models.EncryptionService;
import com.azure.resourcemanager.storage.models.EncryptionServices;
import com.azure.resourcemanager.storage.models.Identity;
import com.azure.resourcemanager.storage.models.IdentityType;
import com.azure.resourcemanager.storage.models.KeySource;
import com.azure.resourcemanager.storage.models.KeyType;
import com.azure.resourcemanager.storage.models.KeyVaultProperties;
import com.azure.resourcemanager.storage.models.Kind;
import com.azure.resourcemanager.storage.models.Sku;
import com.azure.resourcemanager.storage.models.SkuName;
import com.azure.resourcemanager.storage.models.StorageAccountUpdateParameters;
import com.azure.resourcemanager.storage.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StorageAccounts Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/
     * StorageAccountUpdateUserAssignedIdentityWithFederatedIdentityClientId.json
     */
    /**
     * Sample code: StorageAccountUpdateUserAssignedIdentityWithFederatedIdentityClientId.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountUpdateUserAssignedIdentityWithFederatedIdentityClientId(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageAccounts().updateWithResponse("res131918",
            "sto131918",
            new StorageAccountUpdateParameters().withSku(new Sku().withName(SkuName.STANDARD_LRS))
                .withIdentity(new Identity().withType(IdentityType.USER_ASSIGNED).withUserAssignedIdentities(mapOf(
                    "/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}",
                    new UserAssignedIdentity())))
                .withKind(Kind.STORAGE)
                .withEncryption(new Encryption()
                    .withServices(new EncryptionServices()
                        .withBlob(new EncryptionService().withEnabled(true).withKeyType(KeyType.ACCOUNT))
                        .withFile(new EncryptionService().withEnabled(true).withKeyType(KeyType.ACCOUNT)))
                    .withKeySource(KeySource.MICROSOFT_KEYVAULT)
                    .withKeyVaultProperties(new KeyVaultProperties().withKeyName("fakeTokenPlaceholder")
                        .withKeyVersion("fakeTokenPlaceholder").withKeyVaultUri("fakeTokenPlaceholder"))
                    .withEncryptionIdentity(new EncryptionIdentity().withEncryptionUserAssignedIdentity(
                        "/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}")
                        .withEncryptionFederatedIdentityClientId("3109d1c4-a5de-4d84-8832-feabb916a4b6"))),
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
