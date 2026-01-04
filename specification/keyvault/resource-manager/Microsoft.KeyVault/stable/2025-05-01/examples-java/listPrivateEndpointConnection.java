
/**
 * Samples for PrivateEndpointConnections ListByResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/listPrivateEndpointConnection.json
     */
    /**
     * Sample code: KeyVaultListPrivateEndpointConnection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void keyVaultListPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getPrivateEndpointConnections().listByResource("sample-group",
            "sample-vault", com.azure.core.util.Context.NONE);
    }
}
