Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.CheckPrivateLinkServiceVisibilityRequest;

/** Samples for PrivateLinkServices CheckPrivateLinkServiceVisibility. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/CheckPrivateLinkServiceVisibility.json
     */
    /**
     * Sample code: Check private link service visibility.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void checkPrivateLinkServiceVisibility(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPrivateLinkServices()
            .checkPrivateLinkServiceVisibility(
                "westus",
                new CheckPrivateLinkServiceVisibilityRequest()
                    .withPrivateLinkServiceAlias("mypls.00000000-0000-0000-0000-000000000000.azure.privatelinkservice"),
                Context.NONE);
    }
}
```
