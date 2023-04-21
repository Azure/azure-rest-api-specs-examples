/** Samples for ManagedHsms ListDeleted. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-02-01/examples/DeletedManagedHsm_List.json
     */
    /**
     * Sample code: List deleted managed HSMs in the specified subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDeletedManagedHSMsInTheSpecifiedSubscription(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getManagedHsms().listDeleted(com.azure.core.util.Context.NONE);
    }
}
