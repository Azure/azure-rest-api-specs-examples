import com.azure.core.util.Context;

/** Samples for VirtualMachines GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/GetVirtualMachineAutoPlacedOnDedicatedHostGroup.json
     */
    /**
     * Sample code: Get a virtual machine placed on a dedicated host group through automatic placement.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAVirtualMachinePlacedOnADedicatedHostGroupThroughAutomaticPlacement(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .getByResourceGroupWithResponse("myResourceGroup", "myVM", null, Context.NONE);
    }
}
