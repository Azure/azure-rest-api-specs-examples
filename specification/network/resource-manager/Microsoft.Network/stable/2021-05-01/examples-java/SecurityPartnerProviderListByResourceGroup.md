Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SecurityPartnerProviders ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/SecurityPartnerProviderListByResourceGroup.json
     */
    /**
     * Sample code: List all Security Partner Providers for a given resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllSecurityPartnerProvidersForAGivenResourceGroup(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getSecurityPartnerProviders()
            .listByResourceGroup("rg1", Context.NONE);
    }
}
```
