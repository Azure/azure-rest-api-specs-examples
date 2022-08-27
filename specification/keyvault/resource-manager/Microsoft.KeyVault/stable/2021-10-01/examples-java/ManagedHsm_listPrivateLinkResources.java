import com.azure.core.util.Context;

/** Samples for MhsmPrivateLinkResources ListByMhsmResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/ManagedHsm_listPrivateLinkResources.json
     */
    /**
     * Sample code: KeyVaultListPrivateLinkResources.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void keyVaultListPrivateLinkResources(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .vaults()
            .manager()
            .serviceClient()
            .getMhsmPrivateLinkResources()
            .listByMhsmResourceWithResponse("sample-group", "sample-mhsm", Context.NONE);
    }
}
