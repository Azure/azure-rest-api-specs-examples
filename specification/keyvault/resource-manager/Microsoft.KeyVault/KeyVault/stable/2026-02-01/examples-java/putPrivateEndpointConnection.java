
import com.azure.resourcemanager.keyvault.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.keyvault.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.keyvault.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections Put.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/putPrivateEndpointConnection.json
     */
    /**
     * Sample code: KeyVaultPutPrivateEndpointConnection.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void
        keyVaultPutPrivateEndpointConnection(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().putWithResponse("sample-group", "sample-vault",
            "sample-pec",
            new PrivateEndpointConnectionInner().withEtag("").withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState().withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("My name is Joe and I'm approving this.")),
            com.azure.core.util.Context.NONE);
    }
}
