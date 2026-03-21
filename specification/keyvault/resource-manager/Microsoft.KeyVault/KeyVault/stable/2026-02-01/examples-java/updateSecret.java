
import com.azure.resourcemanager.keyvault.models.SecretPatchParameters;
import com.azure.resourcemanager.keyvault.models.SecretPatchProperties;

/**
 * Samples for Secrets Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/updateSecret.json
     */
    /**
     * Sample code: Update a secret.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void updateASecret(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getSecrets().updateWithResponse("sample-group", "sample-vault", "secret-name",
            new SecretPatchParameters().withProperties(new SecretPatchProperties().withValue("secret-value2")),
            com.azure.core.util.Context.NONE);
    }
}
