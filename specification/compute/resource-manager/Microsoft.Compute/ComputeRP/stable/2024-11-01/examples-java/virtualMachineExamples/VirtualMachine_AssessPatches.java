
/**
 * Samples for VirtualMachines AssessPatches.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineExamples/VirtualMachine_AssessPatches.json
     */
    /**
     * Sample code: Assess patch state of a virtual machine.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void assessPatchStateOfAVirtualMachine(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().assessPatches("myResourceGroupName",
            "myVMName", com.azure.core.util.Context.NONE);
    }
}
