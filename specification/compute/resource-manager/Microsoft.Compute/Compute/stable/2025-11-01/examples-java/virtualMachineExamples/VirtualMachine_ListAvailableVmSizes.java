
/**
 * Samples for VirtualMachines ListAvailableSizes.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_ListAvailableVmSizes.json
     */
    /**
     * Sample code: Lists all available virtual machine sizes to which the specified virtual machine can be resized.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listsAllAvailableVirtualMachineSizesToWhichTheSpecifiedVirtualMachineCanBeResized(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().listAvailableSizes("myResourceGroup", "myVmName",
            com.azure.core.util.Context.NONE);
    }
}
