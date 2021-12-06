Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-webpubsub_1.0.0-beta.2/sdk/webpubsub/azure-resourcemanager-webpubsub/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for WebPubSubPrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubPrivateEndpointConnections_Delete.json
     */
    /**
     * Sample code: WebPubSubPrivateEndpointConnections_Delete.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubPrivateEndpointConnectionsDelete(
        com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubPrivateEndpointConnections()
            .delete(
                "mywebpubsubservice.1fa229cd-bf3f-47f0-8c49-afb36723997e",
                "myResourceGroup",
                "myWebPubSubService",
                Context.NONE);
    }
}
```
