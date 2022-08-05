import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMRunCommands Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/runCommands/GetVirtualMachineScaleSetVMRunCommands.json
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
