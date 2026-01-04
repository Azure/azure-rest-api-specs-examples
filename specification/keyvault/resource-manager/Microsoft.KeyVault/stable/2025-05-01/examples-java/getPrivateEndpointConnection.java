
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/getPrivateEndpointConnection.json
     */
    /**
     * Sample code: KeyVaultGetPrivateEndpointConnection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void keyVaultGetPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getPrivateEndpointConnections().getWithResponse("sample-group",
            "sample-vault", "sample-pec", com.azure.core.util.Context.NONE);
    }
}
