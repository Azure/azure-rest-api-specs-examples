
/**
 * Samples for Vaults ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/listVaultByResourceGroup.json
     */
    /**
     * Sample code: List vaults in the specified resource group.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void
        listVaultsInTheSpecifiedResourceGroup(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getVaults().listByResourceGroup("sample-group", 1, com.azure.core.util.Context.NONE);
    }
}
