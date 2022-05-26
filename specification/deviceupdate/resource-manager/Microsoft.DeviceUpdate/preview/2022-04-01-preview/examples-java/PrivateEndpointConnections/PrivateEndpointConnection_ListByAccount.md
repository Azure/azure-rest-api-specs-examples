Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-deviceupdate_1.0.0-beta.1/sdk/deviceupdate/azure-resourcemanager-deviceupdate/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections ListByAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/PrivateEndpointConnections/PrivateEndpointConnection_ListByAccount.json
     */
    /**
     * Sample code: PrivateEndpointConnectionList.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void privateEndpointConnectionList(
        com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.privateEndpointConnections().listByAccount("test-rg", "contoso", Context.NONE);
    }
}
```
