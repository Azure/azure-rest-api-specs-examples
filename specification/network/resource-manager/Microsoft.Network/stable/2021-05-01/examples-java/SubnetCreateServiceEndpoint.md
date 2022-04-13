Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import com.azure.resourcemanager.network.models.ServiceEndpointPropertiesFormat;
import java.util.Arrays;

/** Samples for Subnets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/SubnetCreateServiceEndpoint.json
     */
    /**
     * Sample code: Create subnet with service endpoints.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createSubnetWithServiceEndpoints(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getSubnets()
            .createOrUpdate(
                "subnet-test",
                "vnetname",
                "subnet1",
                new SubnetInner()
                    .withAddressPrefix("10.0.0.0/16")
                    .withServiceEndpoints(
                        Arrays.asList(new ServiceEndpointPropertiesFormat().withService("Microsoft.Storage"))),
                Context.NONE);
    }
}
```
