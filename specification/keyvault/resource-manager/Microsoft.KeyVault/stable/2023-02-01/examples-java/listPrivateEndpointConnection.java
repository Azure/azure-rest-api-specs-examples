/** Samples for PrivateEndpointConnections ListByResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-02-01/examples/listPrivateEndpointConnection.json
     */
    /**
     * Sample code: KeyVaultListPrivateEndpointConnection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void keyVaultListPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .vaults()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .listByResource("sample-group", "sample-vault", com.azure.core.util.Context.NONE);
    }
}
