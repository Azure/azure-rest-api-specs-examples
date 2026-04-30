
/**
 * Samples for VirtualMachines Redeploy.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_Redeploy_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_Redeploy_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineRedeployMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().redeploy("rgcompute", "aaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
