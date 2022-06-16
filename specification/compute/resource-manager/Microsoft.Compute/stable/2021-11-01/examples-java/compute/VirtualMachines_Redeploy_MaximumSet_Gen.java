import com.azure.core.util.Context;

/** Samples for VirtualMachines Redeploy. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachines_Redeploy_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachines_Redeploy_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachinesRedeployMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().redeploy("rgcompute", "a", Context.NONE);
    }
}
