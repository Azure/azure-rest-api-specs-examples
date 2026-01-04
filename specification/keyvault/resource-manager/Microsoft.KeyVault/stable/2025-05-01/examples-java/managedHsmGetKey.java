
/**
 * Samples for ManagedHsmKeys Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/managedHsmGetKey.json
     */
    /**
     * Sample code: Get a key.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getManagedHsmKeys().getWithResponse("sample-group",
            "sample-managedhsm-name", "sample-key-name", com.azure.core.util.Context.NONE);
    }
}
