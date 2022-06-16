import com.azure.core.util.Context;

/** Samples for VirtualMachineSizes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/computeRPCommonExamples/VirtualMachineSizes_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineSizes_List_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineSizesListMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineSizes().list("-e", Context.NONE);
    }
}
