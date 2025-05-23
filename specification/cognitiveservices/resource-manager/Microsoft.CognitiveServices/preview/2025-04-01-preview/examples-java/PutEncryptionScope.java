
import com.azure.resourcemanager.cognitiveservices.models.EncryptionScopeProperties;
import com.azure.resourcemanager.cognitiveservices.models.EncryptionScopeState;
import com.azure.resourcemanager.cognitiveservices.models.KeySource;
import com.azure.resourcemanager.cognitiveservices.models.KeyVaultProperties;

/**
 * Samples for EncryptionScopes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * PutEncryptionScope.json
     */
    /**
     * Sample code: PutEncryptionScope.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        putEncryptionScope(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.encryptionScopes().define("encryptionScopeName").withExistingAccount("resourceGroupName", "accountName")
            .withProperties(new EncryptionScopeProperties()
                .withKeyVaultProperties(new KeyVaultProperties().withKeyName("fakeTokenPlaceholder")
                    .withKeyVersion("fakeTokenPlaceholder").withKeyVaultUri("fakeTokenPlaceholder")
                    .withIdentityClientId("00000000-0000-0000-0000-000000000000"))
                .withKeySource(KeySource.MICROSOFT_KEY_VAULT).withState(EncryptionScopeState.ENABLED))
            .create();
    }
}
