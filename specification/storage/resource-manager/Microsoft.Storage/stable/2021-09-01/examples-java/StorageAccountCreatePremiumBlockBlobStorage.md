Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.models.Encryption;
import com.azure.resourcemanager.storage.models.EncryptionService;
import com.azure.resourcemanager.storage.models.EncryptionServices;
import com.azure.resourcemanager.storage.models.KeySource;
import com.azure.resourcemanager.storage.models.KeyType;
import com.azure.resourcemanager.storage.models.Kind;
import com.azure.resourcemanager.storage.models.MinimumTlsVersion;
import com.azure.resourcemanager.storage.models.Sku;
import com.azure.resourcemanager.storage.models.SkuName;
import com.azure.resourcemanager.storage.models.StorageAccountCreateParameters;
import java.util.HashMap;
import java.util.Map;

/** Samples for StorageAccounts Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountCreatePremiumBlockBlobStorage.json
     */
    /**
     * Sample code: StorageAccountCreatePremiumBlockBlobStorage.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountCreatePremiumBlockBlobStorage(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .create(
                "res9101",
                "sto4445",
                new StorageAccountCreateParameters()
                    .withSku(new Sku().withName(SkuName.PREMIUM_LRS))
                    .withKind(Kind.BLOCK_BLOB_STORAGE)
                    .withLocation("eastus")
                    .withTags(mapOf("key1", "value1", "key2", "value2"))
                    .withEncryption(
                        new Encryption()
                            .withServices(
                                new EncryptionServices()
                                    .withBlob(new EncryptionService().withEnabled(true).withKeyType(KeyType.ACCOUNT))
                                    .withFile(new EncryptionService().withEnabled(true).withKeyType(KeyType.ACCOUNT)))
                            .withKeySource(KeySource.MICROSOFT_STORAGE)
                            .withRequireInfrastructureEncryption(false))
                    .withMinimumTlsVersion(MinimumTlsVersion.TLS1_2)
                    .withAllowSharedKeyAccess(true),
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
```
