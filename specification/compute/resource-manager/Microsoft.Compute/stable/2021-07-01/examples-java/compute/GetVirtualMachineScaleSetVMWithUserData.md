Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/GetVirtualMachineScaleSetVMWithUserData.json
     */
    /**
     * Sample code: Get VM scale set VM with UserData.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVMScaleSetVMWithUserData(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .getWithResponse("myResourceGroup", "{vmss-name}", "0", null, Context.NONE);
    }
}
```
