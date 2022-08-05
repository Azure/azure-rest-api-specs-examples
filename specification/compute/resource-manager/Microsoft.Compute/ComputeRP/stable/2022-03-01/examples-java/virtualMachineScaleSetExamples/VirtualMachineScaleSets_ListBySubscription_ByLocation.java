import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSets ListByLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_ListBySubscription_ByLocation.json
     */
    /**
     * Sample code: Lists all the VM scale sets under the specified subscription for the specified location.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listsAllTheVMScaleSetsUnderTheSpecifiedSubscriptionForTheSpecifiedLocation(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .listByLocation("eastus", Context.NONE);
    }
}
