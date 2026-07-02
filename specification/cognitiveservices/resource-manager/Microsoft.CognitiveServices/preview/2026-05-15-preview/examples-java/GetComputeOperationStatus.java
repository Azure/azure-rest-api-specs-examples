
/**
 * Samples for ComputeOperations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/GetComputeOperationStatus.json
     */
    /**
     * Sample code: GetComputeOperationStatus.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getComputeOperationStatus(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.computeOperations().getWithResponse("eastus", "00000000-0000-0000-0000-000000000000",
            com.azure.core.util.Context.NONE);
    }
}
