Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-iothub_1.2.0-beta.1/sdk/iothub/azure-resourcemanager-iothub/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.iothub.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.iothub.models.PrivateEndpointConnectionProperties;
import com.azure.resourcemanager.iothub.models.PrivateLinkServiceConnectionState;
import com.azure.resourcemanager.iothub.models.PrivateLinkServiceConnectionStatus;

/** Samples for PrivateEndpointConnections Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_updateprivateendpointconnection.json
     */
    /**
     * Sample code: PrivateEndpointConnection_Update.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void privateEndpointConnectionUpdate(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .privateEndpointConnections()
            .update(
                "myResourceGroup",
                "testHub",
                "myPrivateEndpointConnection",
                new PrivateEndpointConnectionInner()
                    .withProperties(
                        new PrivateEndpointConnectionProperties()
                            .withPrivateLinkServiceConnectionState(
                                new PrivateLinkServiceConnectionState()
                                    .withStatus(PrivateLinkServiceConnectionStatus.APPROVED)
                                    .withDescription("Approved by johndoe@contoso.com"))),
                Context.NONE);
    }
}
```
