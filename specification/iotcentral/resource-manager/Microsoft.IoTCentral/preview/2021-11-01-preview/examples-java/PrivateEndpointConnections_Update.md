Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-iotcentral_1.1.0-beta.1/sdk/iotcentral/azure-resourcemanager-iotcentral/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.iotcentral.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.iotcentral.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnections Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/PrivateEndpointConnections_Update.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Update.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void privateEndpointConnectionsUpdate(
        com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager
            .privateEndpointConnections()
            .define("myIoTCentralAppEndpoint")
            .withExistingIotApp("resRg", "myIoTCentralApp")
            .withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState()
                    .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("Auto-approved")
                    .withActionsRequired("None"))
            .create();
    }
}
```
