import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMRunCommands List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/runCommands/ListVirtualMachineScaleSetVMRunCommands.json
     */
    /**
     * Sample code: List run commands in Vmss instance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRunCommandsInVmssInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMRunCommands()
            .list("myResourceGroup", "myvmScaleSet", "0", null, Context.NONE);
    }
}
