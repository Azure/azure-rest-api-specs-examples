
/**
 * Samples for ManagedHsmKeys ListVersions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/managedHsmListKeyVersions.json
     */
    /**
     * Sample code: List key versions in the managed HSM.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listKeyVersionsInTheManagedHSM(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getManagedHsmKeys().listVersions("sample-group",
            "sample-managedhsm-name", "sample-key-name", com.azure.core.util.Context.NONE);
    }
}
