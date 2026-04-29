
/**
 * Samples for VirtualMachineScaleSetVMExtensions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtension_List.json
     */
    /**
     * Sample code: List extensions in Vmss instance.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listExtensionsInVmssInstance(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMExtensions().listWithResponse("myResourceGroup",
            "myvmScaleSet", "0", null, com.azure.core.util.Context.NONE);
    }
}
