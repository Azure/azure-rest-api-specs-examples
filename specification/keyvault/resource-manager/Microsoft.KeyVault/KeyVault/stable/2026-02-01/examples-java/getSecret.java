
/**
 * Samples for Secrets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/getSecret.json
     */
    /**
     * Sample code: Get a secret.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void getASecret(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getSecrets().getWithResponse("sample-group", "sample-vault", "secret-name",
            com.azure.core.util.Context.NONE);
    }
}
