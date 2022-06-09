```java
import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMs RetrieveBootDiagnosticsData. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_RetrieveBootDiagnosticsData.json
     */
    /**
     * Sample code: RetrieveBootDiagnosticsData of a virtual machine.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveBootDiagnosticsDataOfAVirtualMachine(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .retrieveBootDiagnosticsDataWithResponse("ResourceGroup", "myvmScaleSet", "0", 60, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
