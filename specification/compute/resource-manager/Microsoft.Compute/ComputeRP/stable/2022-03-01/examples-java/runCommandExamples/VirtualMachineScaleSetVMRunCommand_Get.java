import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMRunCommands Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/runCommandExamples/VirtualMachineScaleSetVMRunCommand_Get.json
     */
    /**
     * Sample code: Get VirtualMachineScaleSet VM run commands.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualMachineScaleSetVMRunCommands(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMRunCommands()
            .getWithResponse("myResourceGroup", "myvmScaleSet", "0", "myRunCommand", null, Context.NONE);
    }
}
