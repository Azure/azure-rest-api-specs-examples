
/**
 * Samples for ManagedComputeCapacities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ListManagedComputeCapacities.json
     */
    /**
     * Sample code: List Managed Compute Capacities.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listManagedComputeCapacities(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.managedComputeCapacities().list("MaaP", null, null, com.azure.core.util.Context.NONE);
    }
}
