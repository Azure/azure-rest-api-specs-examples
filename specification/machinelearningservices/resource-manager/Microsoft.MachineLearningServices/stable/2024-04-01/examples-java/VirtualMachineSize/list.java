
/**
 * Samples for VirtualMachineSizes List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/VirtualMachineSize/list.json
     */
    /**
     * Sample code: List VM Sizes.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listVMSizes(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.virtualMachineSizes().listWithResponse("eastus", com.azure.core.util.Context.NONE);
    }
}
