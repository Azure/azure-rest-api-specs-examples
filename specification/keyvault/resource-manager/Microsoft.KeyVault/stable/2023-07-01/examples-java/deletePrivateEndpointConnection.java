
/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/
     * deletePrivateEndpointConnection.json
     */
    /**
     * Sample code: KeyVaultDeletePrivateEndpointConnection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void keyVaultDeletePrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getPrivateEndpointConnections().delete("sample-group", "sample-vault",
            "sample-pec", com.azure.core.util.Context.NONE);
    }
}
