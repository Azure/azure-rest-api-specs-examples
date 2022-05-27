Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for VirtualMachines GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachine_Get_WithVMSizeProperties.json
     */
    /**
     * Sample code: Get a virtual machine with VM Size Properties.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAVirtualMachineWithVMSizeProperties(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .getByResourceGroupWithResponse("myResourceGroup", "myVM", null, Context.NONE);
    }
}
```
