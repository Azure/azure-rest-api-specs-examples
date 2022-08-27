import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/getPrivateEndpointConnection.json
     */
    /**
     * Sample code: KeyVaultGetPrivateEndpointConnection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void keyVaultGetPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .vaults()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .getWithResponse("sample-group", "sample-vault", "sample-pec", Context.NONE);
    }
}
