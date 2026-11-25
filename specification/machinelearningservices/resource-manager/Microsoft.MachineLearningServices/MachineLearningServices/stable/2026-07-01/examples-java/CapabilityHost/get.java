
/**
 * Samples for CapabilityHosts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/CapabilityHost/get.json
     */
    /**
     * Sample code: Get CapabilityHost.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getCapabilityHost(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.capabilityHosts().getWithResponse("test-rg", "my-aml-workspace", "capabilityHostName",
            com.azure.core.util.Context.NONE);
    }
}
