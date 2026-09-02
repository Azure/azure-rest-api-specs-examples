
import com.azure.resourcemanager.cognitiveservices.models.AccountProperties;
import com.azure.resourcemanager.cognitiveservices.models.CapabilitySettings;
import com.azure.resourcemanager.cognitiveservices.models.Encryption;
import com.azure.resourcemanager.cognitiveservices.models.Identity;
import com.azure.resourcemanager.cognitiveservices.models.KeySource;
import com.azure.resourcemanager.cognitiveservices.models.KeyVaultProperties;
import com.azure.resourcemanager.cognitiveservices.models.ResourceIdentityType;
import com.azure.resourcemanager.cognitiveservices.models.Sku;
import com.azure.resourcemanager.cognitiveservices.models.UserOwnedStorage;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Accounts Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15-preview/CreateAccount.json
     */
    /**
     * Sample code: Create Account.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createAccount(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().define("testCreate1").withExistingResourceGroup("myResourceGroup").withRegion("West US")
            .withProperties(new AccountProperties()
                .withEncryption(new Encryption()
                    .withKeyVaultProperties(new KeyVaultProperties().withKeyName("fakeTokenPlaceholder")
                        .withKeyVersion("fakeTokenPlaceholder").withKeyVaultUri("fakeTokenPlaceholder"))
                    .withKeySource(KeySource.MICROSOFT_KEY_VAULT))
                .withUserOwnedStorage(Arrays.asList(new UserOwnedStorage().withResourceId(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount")))
                .withCapabilitySettings(new CapabilitySettings().withDocumentStore(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.DocumentDB/databaseAccounts/myCosmosAccount")
                    .withVectorStore(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.Search/searchServices/mySearchService")
                    .withBlobStore(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount")))
            .withKind("Emotion").withSku(new Sku().withName("S0"))
            .withIdentity(new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED)).create();
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
