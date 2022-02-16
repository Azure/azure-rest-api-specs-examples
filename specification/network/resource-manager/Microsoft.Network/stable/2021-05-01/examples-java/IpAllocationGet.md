Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for IpAllocations GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/IpAllocationGet.json
     */
    /**
     * Sample code: Get IpAllocation.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getIpAllocation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getIpAllocations()
            .getByResourceGroupWithResponse("rg1", "test-ipallocation", null, Context.NONE);
    }
}
```
