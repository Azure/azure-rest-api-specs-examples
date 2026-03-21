
/**
 * Samples for Vaults GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/getVault.json
     */
    /**
     * Sample code: Retrieve a vault.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void retrieveAVault(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getVaults().getByResourceGroupWithResponse("sample-resource-group", "sample-vault",
            com.azure.core.util.Context.NONE);
    }
}
