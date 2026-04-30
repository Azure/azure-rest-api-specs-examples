
/**
 * Samples for VirtualMachineScaleSetVMRunCommands List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/runCommandExamples/VirtualMachineScaleSetVMRunCommand_List.json
     */
    /**
     * Sample code: List run commands in Vmss instance.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listRunCommandsInVmssInstance(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMRunCommands().list("myResourceGroup", "myvmScaleSet", "0",
            null, com.azure.core.util.Context.NONE);
    }
}
