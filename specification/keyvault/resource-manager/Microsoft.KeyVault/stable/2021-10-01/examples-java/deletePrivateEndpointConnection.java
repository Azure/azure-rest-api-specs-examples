import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/deletePrivateEndpointConnection.json
     */
    /**
     * Sample code: KeyVaultDeletePrivateEndpointConnection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void keyVaultDeletePrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .vaults()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .delete("sample-group", "sample-vault", "sample-pec", Context.NONE);
    }
}
