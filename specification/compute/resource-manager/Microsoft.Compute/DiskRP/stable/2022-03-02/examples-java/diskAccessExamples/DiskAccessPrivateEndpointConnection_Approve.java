import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.compute.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.compute.models.PrivateLinkServiceConnectionState;

/** Samples for DiskAccesses UpdateAPrivateEndpointConnection. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-03-02/examples/diskAccessExamples/DiskAccessPrivateEndpointConnection_Approve.json
     */
    /**
     * Sample code: Approve a Private Endpoint Connection under a disk access resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void approveAPrivateEndpointConnectionUnderADiskAccessResource(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDiskAccesses()
            .updateAPrivateEndpointConnection(
                "myResourceGroup",
                "myDiskAccess",
                "myPrivateEndpointConnection",
                new PrivateEndpointConnectionInner()
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkServiceConnectionState()
                            .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                            .withDescription("Approving myPrivateEndpointConnection")),
                Context.NONE);
    }
}
