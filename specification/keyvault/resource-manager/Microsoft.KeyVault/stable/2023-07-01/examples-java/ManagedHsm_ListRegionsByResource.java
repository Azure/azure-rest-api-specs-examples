
/**
 * Samples for MhsmRegions ListByResource.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/
     * ManagedHsm_ListRegionsByResource.json
     */
    /**
     * Sample code: List managed HSM Pools in a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagedHSMPoolsInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getMhsmRegions().listByResource("sample-group", "sample-mhsm",
            com.azure.core.util.Context.NONE);
    }
}
