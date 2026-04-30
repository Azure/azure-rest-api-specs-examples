
/**
 * Samples for VirtualMachineScaleSetVMRunCommands Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/runCommandExamples/VirtualMachineScaleSetVMRunCommand_Delete.json
     */
    /**
     * Sample code: Delete VirtualMachineScaleSet VM run command.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        deleteVirtualMachineScaleSetVMRunCommand(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMRunCommands().delete("myResourceGroup", "myvmScaleSet", "0",
            "myRunCommand", com.azure.core.util.Context.NONE);
    }
}
