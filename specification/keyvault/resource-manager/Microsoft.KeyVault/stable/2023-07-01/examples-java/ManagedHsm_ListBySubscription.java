
/** Samples for ManagedHsms List. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/
     * ManagedHsm_ListBySubscription.json
     */
    /**
     * Sample code: List managed HSM Pools in a subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagedHSMPoolsInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getManagedHsms().list(null, com.azure.core.util.Context.NONE);
    }
}
