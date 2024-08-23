
/**
 * Samples for ManagedHsmKeys Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/managedHsmGetKey.json
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
