Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.storage.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.storage.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnections Put. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/StorageAccountPutPrivateEndpointConnection.json
     */
    /**
     * Sample code: StorageAccountPutPrivateEndpointConnection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountPutPrivateEndpointConnection(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getPrivateEndpointConnections()
            .putWithResponse(
                "res7687",
                "sto9699",
                "{privateEndpointConnectionName}",
                new PrivateEndpointConnectionInner()
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkServiceConnectionState()
                            .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                            .withDescription("Auto-Approved")),
                Context.NONE);
    }
}
```
