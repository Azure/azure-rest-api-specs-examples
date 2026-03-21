
/**
 * Samples for ManagedHsmKeys ListVersions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/managedHsmListKeyVersions.json
     */
    /**
     * Sample code: List key versions in the managed HSM.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void listKeyVersionsInTheManagedHSM(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getManagedHsmKeys().listVersions("sample-group", "sample-managedhsm-name",
            "sample-key-name", com.azure.core.util.Context.NONE);
    }
}
