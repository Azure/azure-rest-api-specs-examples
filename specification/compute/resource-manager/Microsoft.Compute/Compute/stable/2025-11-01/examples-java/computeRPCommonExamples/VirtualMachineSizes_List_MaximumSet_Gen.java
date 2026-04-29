
/**
 * Samples for VirtualMachineSizes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/computeRPCommonExamples/VirtualMachineSizes_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineSizes_List_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineSizesListMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineSizes().list("-e", com.azure.core.util.Context.NONE);
    }
}
