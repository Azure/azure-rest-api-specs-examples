import com.azure.core.util.Context;

/** Samples for PrivateLinkResources ListByVault. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/listPrivateLinkResources.json
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
            .getPrivateLinkResources()
            .listByVaultWithResponse("sample-group", "sample-vault", Context.NONE);
    }
}
