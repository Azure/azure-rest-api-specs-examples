
/**
 * Samples for VirtualMachineSizes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/computeRPCommonExamples/VirtualMachineSizes_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineSizes_List_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineSizesListMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineSizes().list("._..", com.azure.core.util.Context.NONE);
    }
}
