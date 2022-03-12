Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.compute.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.compute.models.PrivateLinkServiceConnectionState;

/** Samples for DiskAccesses UpdateAPrivateEndpointConnection. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-12-01/examples/ApprovePrivateEndpointConnection.json
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
```
