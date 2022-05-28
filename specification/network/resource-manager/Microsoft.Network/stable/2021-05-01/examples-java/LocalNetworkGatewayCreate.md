Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.LocalNetworkGatewayInner;
import com.azure.resourcemanager.network.models.AddressSpace;
import java.util.Arrays;

/** Samples for LocalNetworkGateways CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LocalNetworkGatewayCreate.json
     */
    /**
     * Sample code: CreateLocalNetworkGateway.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createLocalNetworkGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getLocalNetworkGateways()
            .createOrUpdate(
                "rg1",
                "localgw",
                new LocalNetworkGatewayInner()
                    .withLocation("Central US")
                    .withLocalNetworkAddressSpace(new AddressSpace().withAddressPrefixes(Arrays.asList("10.1.0.0/16")))
                    .withGatewayIpAddress("11.12.13.14")
                    .withFqdn("site1.contoso.com"),
                Context.NONE);
    }
}
```
