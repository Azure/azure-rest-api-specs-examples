
/**
 * Samples for Secrets Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/getSecret.json
     */
    /**
     * Sample code: Get a secret.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getASecret(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getSecrets().getWithResponse("sample-group", "sample-vault",
            "secret-name", com.azure.core.util.Context.NONE);
    }
}
