
/**
 * Samples for CapabilityHosts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/CapabilityHost/delete.json
     */
    /**
     * Sample code: Delete CapabilityHost.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteCapabilityHost(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.capabilityHosts().delete("test-rg", "my-aml-workspace", "capabilityHostName",
            com.azure.core.util.Context.NONE);
    }
}
