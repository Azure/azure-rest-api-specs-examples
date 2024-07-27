
import com.azure.resourcemanager.storage.models.Encryption;
import com.azure.resourcemanager.storage.models.EncryptionService;
import com.azure.resourcemanager.storage.models.EncryptionServices;
import com.azure.resourcemanager.storage.models.ExpirationAction;
import com.azure.resourcemanager.storage.models.ExtendedLocation;
import com.azure.resourcemanager.storage.models.ExtendedLocationTypes;
import com.azure.resourcemanager.storage.models.KeyPolicy;
import com.azure.resourcemanager.storage.models.KeySource;
import com.azure.resourcemanager.storage.models.KeyType;
import com.azure.resourcemanager.storage.models.Kind;
import com.azure.resourcemanager.storage.models.MinimumTlsVersion;
import com.azure.resourcemanager.storage.models.RoutingChoice;
import com.azure.resourcemanager.storage.models.RoutingPreference;
import com.azure.resourcemanager.storage.models.SasPolicy;
import com.azure.resourcemanager.storage.models.Sku;
import com.azure.resourcemanager.storage.models.SkuName;
import com.azure.resourcemanager.storage.models.StorageAccountCreateParameters;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StorageAccounts Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountCreate.json
     */
    /**
     * Sample code: StorageAccountCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageAccounts()
            .create("res9101", "sto4445",
                new StorageAccountCreateParameters().withSku(new Sku().withName(SkuName.STANDARD_GRS))
                    .withKind(Kind.STORAGE).withLocation("eastus")
                    .withExtendedLocation(
                        new ExtendedLocation().withName("losangeles001").withType(ExtendedLocationTypes.EDGE_ZONE))
                    .withTags(
                        mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder"))
                    .withSasPolicy(
                        new SasPolicy().withSasExpirationPeriod(
                            "1.15:59:59").withExpirationAction(
                                ExpirationAction.LOG))
                    .withKeyPolicy(new KeyPolicy().withKeyExpirationPeriodInDays(20))
                    .withEncryption(new Encryption()
                        .withServices(new EncryptionServices()
                            .withBlob(new EncryptionService().withEnabled(true).withKeyType(KeyType.ACCOUNT))
                            .withFile(new EncryptionService().withEnabled(true).withKeyType(KeyType.ACCOUNT)))
                        .withKeySource(KeySource.MICROSOFT_STORAGE).withRequireInfrastructureEncryption(false))
                    .withIsSftpEnabled(true).withIsHnsEnabled(true)
                    .withRoutingPreference(new RoutingPreference().withRoutingChoice(RoutingChoice.MICROSOFT_ROUTING)
                        .withPublishMicrosoftEndpoints(true).withPublishInternetEndpoints(true))
                    .withAllowBlobPublicAccess(false).withMinimumTlsVersion(MinimumTlsVersion.TLS1_2)
                    .withAllowSharedKeyAccess(true).withDefaultToOAuthAuthentication(false),
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
