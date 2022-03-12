Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DedicatedHosts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetADedicatedHost.json
     */
    /**
     * Sample code: Get a dedicated host.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getADedicatedHost(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDedicatedHosts()
            .getWithResponse("myResourceGroup", "myDedicatedHostGroup", "myHost", null, Context.NONE);
    }
}
```
