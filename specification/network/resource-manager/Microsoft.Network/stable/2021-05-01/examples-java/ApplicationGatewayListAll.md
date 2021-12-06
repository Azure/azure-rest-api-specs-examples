Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ApplicationGateways List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ApplicationGatewayListAll.json
     */
    /**
     * Sample code: Lists all application gateways in a subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listsAllApplicationGatewaysInASubscription(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationGateways().list(Context.NONE);
    }
}
```
