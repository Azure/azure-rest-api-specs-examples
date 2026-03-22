
/**
 * Samples for ManagedHsmKeys GetVersion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/managedHsmGetKeyVersion.json
     */
    /**
     * Sample code: Get a key version.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void getAKeyVersion(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getManagedHsmKeys().getVersionWithResponse("sample-group", "sample-managedhsm-name",
            "sample-key-name", "fd618d9519b74f9aae94ade66b876acc", com.azure.core.util.Context.NONE);
    }
}
