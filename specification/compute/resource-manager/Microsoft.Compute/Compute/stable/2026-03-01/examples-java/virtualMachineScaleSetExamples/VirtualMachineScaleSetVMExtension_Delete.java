
/**
 * Samples for VirtualMachineScaleSetVMExtensions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtension_Delete.json
     */
    /**
     * Sample code: Delete VirtualMachineScaleSet VM extension.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        deleteVirtualMachineScaleSetVMExtension(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMExtensions().delete("myResourceGroup", "myvmScaleSet", "0",
            "myVMExtension", com.azure.core.util.Context.NONE);
    }
}
