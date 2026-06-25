
/**
 * Samples for VirtualMachines Redeploy.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Redeploy_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_Redeploy_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineRedeployMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().redeploy("rgcompute", "a", com.azure.core.util.Context.NONE);
    }
}
