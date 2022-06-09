```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.webpubsub.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.webpubsub.models.PrivateEndpoint;
import com.azure.resourcemanager.webpubsub.models.PrivateLinkServiceConnectionState;
import com.azure.resourcemanager.webpubsub.models.PrivateLinkServiceConnectionStatus;

/** Samples for WebPubSubPrivateEndpointConnections Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubPrivateEndpointConnections_Update.json
     */
    /**
     * Sample code: WebPubSubPrivateEndpointConnections_Update.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubPrivateEndpointConnectionsUpdate(
        com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubPrivateEndpointConnections()
            .updateWithResponse(
                "mywebpubsubservice.1fa229cd-bf3f-47f0-8c49-afb36723997e",
                "myResourceGroup",
                "myWebPubSubService",
                new PrivateEndpointConnectionInner()
                    .withPrivateEndpoint(
                        new PrivateEndpoint()
                            .withId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"))
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkServiceConnectionState()
                            .withStatus(PrivateLinkServiceConnectionStatus.APPROVED)
                            .withActionsRequired("None")),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-webpubsub_1.0.0-beta.2/sdk/webpubsub/azure-resourcemanager-webpubsub/README.md) on how to add the SDK to your project and authenticate.
