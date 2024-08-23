
/**
 * Samples for ManagedHsmKeys ListVersions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/managedHsmListKeyVersions.
     * json
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
