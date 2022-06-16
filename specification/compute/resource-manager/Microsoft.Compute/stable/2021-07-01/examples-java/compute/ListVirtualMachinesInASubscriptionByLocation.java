import com.azure.core.util.Context;

/** Samples for VirtualMachines ListByLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/ListVirtualMachinesInASubscriptionByLocation.json
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
