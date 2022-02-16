Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-digitaltwins_1.0.0-beta.2/sdk/digitaltwins/azure-resourcemanager-digitaltwins/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/PrivateEndpointConnectionDelete_example.json
     */
    /**
     * Sample code: Delete private endpoint connection with the specified name.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void deletePrivateEndpointConnectionWithTheSpecifiedName(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager
            .privateEndpointConnections()
            .delete("resRg", "myDigitalTwinsService", "myPrivateConnection", Context.NONE);
    }
}
```
