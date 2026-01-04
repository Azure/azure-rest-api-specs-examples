
import com.azure.resourcemanager.keyvault.fluent.models.MhsmPrivateEndpointConnectionInner;
import com.azure.resourcemanager.keyvault.models.MhsmPrivateLinkServiceConnectionState;
import com.azure.resourcemanager.keyvault.models.PrivateEndpointServiceConnectionStatus;

/**
 * Samples for MhsmPrivateEndpointConnections Put.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ManagedHsm_putPrivateEndpointConnection.json
     */
    /**
     * Sample code: ManagedHsmPutPrivateEndpointConnection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void managedHsmPutPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getMhsmPrivateEndpointConnections().putWithResponse("sample-group",
            "sample-mhsm", "sample-pec",
            new MhsmPrivateEndpointConnectionInner().withPrivateLinkServiceConnectionState(
                new MhsmPrivateLinkServiceConnectionState().withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("My name is Joe and I'm approving this.")),
            com.azure.core.util.Context.NONE);
    }
}
