
/**
 * Samples for VirtualMachineScaleSets Redeploy.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Redeploy_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_Redeploy_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetRedeployMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().redeploy("rgcompute", "aaaaaaaaaaaaaaaaaaaaaa", null,
            com.azure.core.util.Context.NONE);
    }
}
