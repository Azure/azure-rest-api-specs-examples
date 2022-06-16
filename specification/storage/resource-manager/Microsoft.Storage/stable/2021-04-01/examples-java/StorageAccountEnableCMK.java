import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.models.Encryption;
import com.azure.resourcemanager.storage.models.EncryptionService;
import com.azure.resourcemanager.storage.models.EncryptionServices;
import com.azure.resourcemanager.storage.models.KeySource;
import com.azure.resourcemanager.storage.models.KeyType;
import com.azure.resourcemanager.storage.models.KeyVaultProperties;
import com.azure.resourcemanager.storage.models.StorageAccountUpdateParameters;
import java.util.HashMap;
import java.util.Map;

/** Samples for StorageAccounts Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/StorageAccountEnableCMK.json
     */
    /**
     * Sample code: StorageAccountEnableCMK.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountEnableCMK(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .updateWithResponse(
                "res9407",
                "sto8596",
                new StorageAccountUpdateParameters()
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
                                    .withKeyVaultUri("https://myvault8569.vault.azure.net"))),
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
