import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMExtensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtensions_List.json
     */
    /**
     * Sample code: List extensions in Vmss instance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listExtensionsInVmssInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMExtensions()
            .listWithResponse("myResourceGroup", "myvmScaleSet", "0", null, Context.NONE);
    }
}
