import com.azure.core.util.Context;

/** Samples for VirtualMachines Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachine_Delete_Force.json
     */
    /**
     * Sample code: Force delete a VM.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void forceDeleteAVM(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .delete("myResourceGroup", "myVM", true, Context.NONE);
    }
}
