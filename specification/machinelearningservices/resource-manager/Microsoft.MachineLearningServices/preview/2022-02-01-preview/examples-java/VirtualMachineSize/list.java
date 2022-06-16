import com.azure.core.util.Context;

/** Samples for VirtualMachineSizes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/VirtualMachineSize/list.json
     */
    /**
     * Sample code: List VM Sizes.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listVMSizes(com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.virtualMachineSizes().listWithResponse("eastus", Context.NONE);
    }
}
