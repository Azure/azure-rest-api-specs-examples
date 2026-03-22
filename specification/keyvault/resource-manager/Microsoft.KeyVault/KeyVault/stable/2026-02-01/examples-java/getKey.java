
/**
 * Samples for Keys Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/getKey.json
     */
    /**
     * Sample code: Get a key.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void getAKey(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getKeys().getWithResponse("sample-group", "sample-vault-name", "sample-key-name",
            com.azure.core.util.Context.NONE);
    }
}
