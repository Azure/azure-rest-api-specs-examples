Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for NetworkInterfaces ListVirtualMachineScaleSetVMNetworkInterfaces. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VmssVmNetworkInterfaceList.json
     */
    /**
     * Sample code: List virtual machine scale set vm network interfaces.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVirtualMachineScaleSetVmNetworkInterfaces(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkInterfaces()
            .listVirtualMachineScaleSetVMNetworkInterfaces("rg1", "vmss1", "1", Context.NONE);
    }
}
```
