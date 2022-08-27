import com.azure.core.util.Context;

/** Samples for VirtualMachines ListByLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineExamples/VirtualMachines_ListBySubscription_ByLocation.json
     */
    /**
     * Sample code: Lists all the virtual machines under the specified subscription for the specified location.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listsAllTheVirtualMachinesUnderTheSpecifiedSubscriptionForTheSpecifiedLocation(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().listByLocation("eastus", Context.NONE);
    }
}
