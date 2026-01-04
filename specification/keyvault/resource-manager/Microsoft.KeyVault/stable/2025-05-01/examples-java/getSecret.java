
/**
 * Samples for Secrets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/getSecret.json
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
