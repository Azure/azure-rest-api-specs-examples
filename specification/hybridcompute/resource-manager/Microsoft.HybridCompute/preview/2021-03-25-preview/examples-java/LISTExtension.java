/** Samples for MachineExtensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-03-25-preview/examples/LISTExtension.json
     */
    /**
     * Sample code: GET all Machine Extensions.
     *
     * @param manager Entry point to HybridComputeManager.
     */
    public static void gETAllMachineExtensions(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.machineExtensions().list("myResourceGroup", "myMachine", null, com.azure.core.util.Context.NONE);
    }
}
