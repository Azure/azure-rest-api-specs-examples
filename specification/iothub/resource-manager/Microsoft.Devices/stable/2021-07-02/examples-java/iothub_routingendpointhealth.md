Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-iothub_1.2.0-beta.1/sdk/iothub/azure-resourcemanager-iothub/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for IotHubResource GetEndpointHealth. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_routingendpointhealth.json
     */
    /**
     * Sample code: IotHubResource_GetEndpointHealth.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceGetEndpointHealth(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().getEndpointHealth("myResourceGroup", "testHub", Context.NONE);
    }
}
```
