
/**
 * Samples for Compute Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Compute/get/AmlCompute.json
     */
    /**
     * Sample code: Get a AML Compute.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getAAMLCompute(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.computes().getWithResponse("testrg123", "workspaces123", "compute123",
            com.azure.core.util.Context.NONE);
    }
}
