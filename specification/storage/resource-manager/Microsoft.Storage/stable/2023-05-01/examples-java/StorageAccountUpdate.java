
import com.azure.resourcemanager.storage.models.DefaultAction;
import com.azure.resourcemanager.storage.models.Encryption;
import com.azure.resourcemanager.storage.models.EncryptionService;
import com.azure.resourcemanager.storage.models.EncryptionServices;
import com.azure.resourcemanager.storage.models.ExpirationAction;
import com.azure.resourcemanager.storage.models.KeyPolicy;
import com.azure.resourcemanager.storage.models.KeySource;
import com.azure.resourcemanager.storage.models.KeyType;
import com.azure.resourcemanager.storage.models.MinimumTlsVersion;
import com.azure.resourcemanager.storage.models.NetworkRuleSet;
import com.azure.resourcemanager.storage.models.ResourceAccessRule;
import com.azure.resourcemanager.storage.models.RoutingChoice;
import com.azure.resourcemanager.storage.models.RoutingPreference;
import com.azure.resourcemanager.storage.models.SasPolicy;
import com.azure.resourcemanager.storage.models.StorageAccountUpdateParameters;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StorageAccounts Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountUpdate.json
     */
    /**
     * Sample code: StorageAccountUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageAccounts().updateWithResponse("res9407", "sto8596",
            new StorageAccountUpdateParameters()
                .withEncryption(new Encryption()
                    .withServices(new EncryptionServices()
                        .withBlob(new EncryptionService().withEnabled(true).withKeyType(KeyType.ACCOUNT))
                        .withFile(new EncryptionService().withEnabled(true).withKeyType(KeyType.ACCOUNT)))
                    .withKeySource(KeySource.MICROSOFT_STORAGE))
                .withSasPolicy(
                    new SasPolicy().withSasExpirationPeriod("1.15:59:59").withExpirationAction(ExpirationAction.LOG))
                .withKeyPolicy(new KeyPolicy().withKeyExpirationPeriodInDays(20)).withIsSftpEnabled(true)
                .withIsLocalUserEnabled(true).withEnableExtendedGroups(true)
                .withNetworkRuleSet(new NetworkRuleSet().withResourceAccessRules(Arrays.asList(
                    new ResourceAccessRule().withTenantId("72f988bf-86f1-41af-91ab-2d7cd011db47").withResourceId(
                        "/subscriptions/a7e99807-abbf-4642-bdec-2c809a96a8bc/resourceGroups/res9407/providers/Microsoft.Synapse/workspaces/testworkspace")))
                    .withDefaultAction(DefaultAction.ALLOW))
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
