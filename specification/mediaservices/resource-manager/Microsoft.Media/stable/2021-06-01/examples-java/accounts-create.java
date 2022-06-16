import com.azure.resourcemanager.mediaservices.models.AccountEncryption;
import com.azure.resourcemanager.mediaservices.models.AccountEncryptionKeyType;
import com.azure.resourcemanager.mediaservices.models.KeyVaultProperties;
import com.azure.resourcemanager.mediaservices.models.MediaServiceIdentity;
import com.azure.resourcemanager.mediaservices.models.ResourceIdentity;
import com.azure.resourcemanager.mediaservices.models.StorageAccount;
import com.azure.resourcemanager.mediaservices.models.StorageAccountType;
import com.azure.resourcemanager.mediaservices.models.StorageAuthentication;
import com.azure.resourcemanager.mediaservices.models.UserAssignedManagedIdentity;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Mediaservices CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/accounts-create.json
     */
    /**
     * Sample code: Create a Media Services account.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createAMediaServicesAccount(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .mediaservices()
            .define("contososports")
            .withRegion("South Central US")
            .withExistingResourceGroup("contoso")
            .withTags(mapOf("key1", "value1", "key2", "value2"))
            .withIdentity(
                new MediaServiceIdentity()
                    .withType("UserAssigned")
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1",
                            new UserAssignedManagedIdentity(),
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2",
                            new UserAssignedManagedIdentity())))
            .withStorageAccounts(
                Arrays
                    .asList(
                        new StorageAccount()
                            .withId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Storage/storageAccounts/contososportsstore")
                            .withType(StorageAccountType.PRIMARY)
                            .withIdentity(
                                new ResourceIdentity()
                                    .withUserAssignedIdentity(
                                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1")
                                    .withUseSystemAssignedIdentity(false))))
            .withStorageAuthentication(StorageAuthentication.MANAGED_IDENTITY)
            .withEncryption(
                new AccountEncryption()
                    .withType(AccountEncryptionKeyType.CUSTOMER_KEY)
                    .withKeyVaultProperties(
                        new KeyVaultProperties().withKeyIdentifier("https://keyvault.vault.azure.net/keys/key1"))
                    .withIdentity(
                        new ResourceIdentity()
                            .withUserAssignedIdentity(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1")
                            .withUseSystemAssignedIdentity(false)))
            .create();
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
