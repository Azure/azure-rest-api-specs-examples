
import com.azure.resourcemanager.keyvault.models.SecretPatchParameters;
import com.azure.resourcemanager.keyvault.models.SecretPatchProperties;

/**
 * Samples for Secrets Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/updateSecret.json
     */
    /**
     * Sample code: Update a secret.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateASecret(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getSecrets().updateWithResponse("sample-group", "sample-vault",
            "secret-name",
            new SecretPatchParameters().withProperties(new SecretPatchProperties().withValue("secret-value2")),
            com.azure.core.util.Context.NONE);
    }
}
