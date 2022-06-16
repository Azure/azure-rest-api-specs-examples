import com.azure.core.util.Context;

/** Samples for Vaults ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2019-09-01/examples/listVaultByResourceGroup.json
     */
    /**
     * Sample code: List vaults in the specified resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVaultsInTheSpecifiedResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getVaults().listByResourceGroup("sample-group", 1, Context.NONE);
    }
}
