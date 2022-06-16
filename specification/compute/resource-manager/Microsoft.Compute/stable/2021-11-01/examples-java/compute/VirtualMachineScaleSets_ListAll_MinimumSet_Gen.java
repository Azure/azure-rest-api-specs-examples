import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSets List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_ListAll_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_ListAll_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetsListAllMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets().list(Context.NONE);
    }
}
