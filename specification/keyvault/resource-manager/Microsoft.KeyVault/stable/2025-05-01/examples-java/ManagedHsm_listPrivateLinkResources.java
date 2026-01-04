
/**
 * Samples for MhsmPrivateLinkResources ListByMhsmResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ManagedHsm_listPrivateLinkResources.json
     */
    /**
     * Sample code: KeyVaultListPrivateLinkResources.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void keyVaultListPrivateLinkResources(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getMhsmPrivateLinkResources()
            .listByMhsmResourceWithResponse("sample-group", "sample-mhsm", com.azure.core.util.Context.NONE);
    }
}
