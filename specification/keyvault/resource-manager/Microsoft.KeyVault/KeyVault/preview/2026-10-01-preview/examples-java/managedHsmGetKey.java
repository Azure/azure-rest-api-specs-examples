
/**
 * Samples for ManagedHsmKeys Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/managedHsmGetKey.json
     */
    /**
     * Sample code: Get a key.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void getAKey(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getManagedHsmKeys().getWithResponse("sample-group", "sample-managedhsm-name",
            "sample-key-name", com.azure.core.util.Context.NONE);
    }
}
