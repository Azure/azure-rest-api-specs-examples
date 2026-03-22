
/**
 * Samples for ManagedHsmKeys List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/managedHsmListKeys.json
     */
    /**
     * Sample code: List keys in the managed HSM.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void listKeysInTheManagedHSM(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getManagedHsmKeys().list("sample-group", "sample-managedhsm-name",
            com.azure.core.util.Context.NONE);
    }
}
