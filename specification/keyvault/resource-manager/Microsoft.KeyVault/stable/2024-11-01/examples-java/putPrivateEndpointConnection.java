
import com.azure.resourcemanager.keyvault.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.keyvault.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.keyvault.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections Put.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/
     * putPrivateEndpointConnection.json
     */
    /**
     * Sample code: KeyVaultPutPrivateEndpointConnection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void keyVaultPutPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getPrivateEndpointConnections().putWithResponse("sample-group",
            "sample-vault", "sample-pec",
            new PrivateEndpointConnectionInner().withEtag("").withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState().withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("My name is Joe and I'm approving this.")),
            com.azure.core.util.Context.NONE);
    }
}
