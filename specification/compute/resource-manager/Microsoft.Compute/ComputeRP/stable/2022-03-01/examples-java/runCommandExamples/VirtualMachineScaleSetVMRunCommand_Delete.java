import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMRunCommands Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/runCommandExamples/VirtualMachineScaleSetVMRunCommand_Delete.json
     */
    /**
     * Sample code: Delete VirtualMachineScaleSet VM run command.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteVirtualMachineScaleSetVMRunCommand(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMRunCommands()
            .delete("myResourceGroup", "myvmScaleSet", "0", "myRunCommand", Context.NONE);
    }
}
