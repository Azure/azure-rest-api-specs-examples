```java
import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/private-endpoint-connection-list.json
     */
    /**
     * Sample code: Get all private endpoint connections.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getAllPrivateEndpointConnections(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.privateEndpointConnections().listWithResponse("contoso", "contososports", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.
