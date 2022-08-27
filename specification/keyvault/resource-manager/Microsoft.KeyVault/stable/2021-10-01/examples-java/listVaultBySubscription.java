import com.azure.core.util.Context;

/** Samples for Vaults ListBySubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/listVaultBySubscription.json
     */
    /**
     * Sample code: List vaults in the specified subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVaultsInTheSpecifiedSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getVaults().listBySubscription(1, Context.NONE);
    }
}
