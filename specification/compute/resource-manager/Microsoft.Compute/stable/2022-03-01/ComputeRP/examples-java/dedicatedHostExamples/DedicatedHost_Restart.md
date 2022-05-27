Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DedicatedHosts Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/dedicatedHostExamples/DedicatedHost_Restart.json
     */
    /**
     * Sample code: Restart Dedicated Host.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void restartDedicatedHost(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDedicatedHosts()
            .restart("myResourceGroup", "myDedicatedHostGroup", "myHost", Context.NONE);
    }
}
```
