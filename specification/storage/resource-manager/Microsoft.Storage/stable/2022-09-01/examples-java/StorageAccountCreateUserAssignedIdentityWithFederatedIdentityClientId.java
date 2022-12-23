import com.azure.core.util.Context;
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
import com.azure.resourcemanager.storage.models.StorageAccountCreateParameters;
import com.azure.resourcemanager.storage.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/** Samples for StorageAccounts Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountCreateUserAssignedIdentityWithFederatedIdentityClientId.json
     */
    /**
     * Sample code: StorageAccountCreateUserAssignedIdentityWithFederatedIdentityClientId.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountCreateUserAssignedIdentityWithFederatedIdentityClientId(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .create(
                "res131918",
                "sto131918",
                new StorageAccountCreateParameters()
                    .withSku(new Sku().withName(SkuName.STANDARD_LRS))
                    .withKind(Kind.STORAGE)
                    .withLocation("eastus")
                    .withIdentity(
                        new Identity()
                            .withType(IdentityType.USER_ASSIGNED)
                            .withUserAssignedIdentities(
                                mapOf(
                                    "/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}",
                                    new UserAssignedIdentity())))
                    .withEncryption(
                        new Encryption()
                            .withServices(
                                new EncryptionServices()
                                    .withBlob(new EncryptionService().withEnabled(true).withKeyType(KeyType.ACCOUNT))
                                    .withFile(new EncryptionService().withEnabled(true).withKeyType(KeyType.ACCOUNT)))
                            .withKeySource(KeySource.MICROSOFT_KEYVAULT)
                            .withKeyVaultProperties(
                                new KeyVaultProperties()
                                    .withKeyName("wrappingKey")
                                    .withKeyVersion("")
                                    .withKeyVaultUri("https://myvault8569.vault.azure.net"))
                            .withEncryptionIdentity(
                                new EncryptionIdentity()
                                    .withEncryptionUserAssignedIdentity(
                                        "/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}")
                                    .withEncryptionFederatedIdentityClientId("f83c6b1b-4d34-47e4-bb34-9d83df58b540"))),
                Context.NONE);
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
