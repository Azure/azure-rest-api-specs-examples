```java
import com.azure.core.util.Context;

/** Samples for ApplicationGateways ListAvailableResponseHeaders. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ApplicationGatewayAvailableResponseHeadersGet.json
     */
    /**
     * Sample code: Get Available Response Headers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAvailableResponseHeaders(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getApplicationGateways()
            .listAvailableResponseHeadersWithResponse(Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
