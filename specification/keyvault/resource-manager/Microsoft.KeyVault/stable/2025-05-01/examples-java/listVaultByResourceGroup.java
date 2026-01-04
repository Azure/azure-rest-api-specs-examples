
/**
 * Samples for Vaults ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/listVaultByResourceGroup.json
     */
    /**
     * Sample code: List vaults in the specified resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVaultsInTheSpecifiedResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getVaults().listByResourceGroup("sample-group", 1,
            com.azure.core.util.Context.NONE);
    }
}
