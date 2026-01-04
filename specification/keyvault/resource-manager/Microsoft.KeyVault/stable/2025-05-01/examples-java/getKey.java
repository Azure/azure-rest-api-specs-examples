
/**
 * Samples for Keys Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/getKey.json
     */
    /**
     * Sample code: Get a key.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getKeys().getWithResponse("sample-group", "sample-vault-name",
            "sample-key-name", com.azure.core.util.Context.NONE);
    }
}
