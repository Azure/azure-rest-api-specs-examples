import com.azure.resourcemanager.cognitiveservices.models.AccountProperties;
import com.azure.resourcemanager.cognitiveservices.models.Encryption;
import com.azure.resourcemanager.cognitiveservices.models.Identity;
import com.azure.resourcemanager.cognitiveservices.models.KeySource;
import com.azure.resourcemanager.cognitiveservices.models.KeyVaultProperties;
import com.azure.resourcemanager.cognitiveservices.models.ResourceIdentityType;
import com.azure.resourcemanager.cognitiveservices.models.Sku;
import com.azure.resourcemanager.cognitiveservices.models.UserOwnedStorage;
import java.util.Arrays;

/** Samples for Accounts Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/CreateAccount.json
     */
    /**
     * Sample code: Create Account.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createAccount(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager
            .accounts()
            .define("testCreate1")
            .withExistingResourceGroup("myResourceGroup")
            .withRegion("West US")
            .withKind("Emotion")
            .withSku(new Sku().withName("S0"))
            .withIdentity(new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
            .withProperties(
                new AccountProperties()
                    .withEncryption(
                        new Encryption()
                            .withKeyVaultProperties(
                                new KeyVaultProperties()
                                    .withKeyName("KeyName")
                                    .withKeyVersion("891CF236-D241-4738-9462-D506AF493DFA")
                                    .withKeyVaultUri("https://pltfrmscrts-use-pc-dev.vault.azure.net/"))
                            .withKeySource(KeySource.MICROSOFT_KEY_VAULT))
                    .withUserOwnedStorage(
                        Arrays
                            .asList(
                                new UserOwnedStorage()
                                    .withResourceId(
                                        "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount"))))
            .create();
    }
}
