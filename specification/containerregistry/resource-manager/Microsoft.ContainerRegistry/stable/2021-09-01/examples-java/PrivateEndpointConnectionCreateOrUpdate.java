import com.azure.core.util.Context;
import com.azure.resourcemanager.containerregistry.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.containerregistry.models.ConnectionStatus;
import com.azure.resourcemanager.containerregistry.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2021-09-01/examples/PrivateEndpointConnectionCreateOrUpdate.json
     */
    /**
     * Sample code: PrivateEndpointConnectionCreateOrUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void privateEndpointConnectionCreateOrUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .createOrUpdate(
                "myResourceGroup",
                "myRegistry",
                "myConnection",
                new PrivateEndpointConnectionInner()
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkServiceConnectionState()
                            .withStatus(ConnectionStatus.APPROVED)
                            .withDescription("Auto-Approved")),
                Context.NONE);
    }
}
