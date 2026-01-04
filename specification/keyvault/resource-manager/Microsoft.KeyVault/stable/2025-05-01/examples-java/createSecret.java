
import com.azure.resourcemanager.keyvault.models.SecretCreateOrUpdateParameters;
import com.azure.resourcemanager.keyvault.models.SecretProperties;

/**
 * Samples for Secrets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/createSecret.json
     */
    /**
     * Sample code: Create a secret.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createASecret(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getSecrets().createOrUpdateWithResponse("sample-group", "sample-vault",
            "secret-name",
            new SecretCreateOrUpdateParameters().withProperties(new SecretProperties().withValue("secret-value")),
            com.azure.core.util.Context.NONE);
    }
}
