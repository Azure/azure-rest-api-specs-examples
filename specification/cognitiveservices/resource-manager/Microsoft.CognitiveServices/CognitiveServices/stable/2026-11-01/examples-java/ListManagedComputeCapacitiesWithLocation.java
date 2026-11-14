
/**
 * Samples for ManagedComputeCapacities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ListManagedComputeCapacitiesWithLocation.json
     */
    /**
     * Sample code: List Managed Compute Capacities With Location.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listManagedComputeCapacitiesWithLocation(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.managedComputeCapacities().list("MaaP", "H100_80GB", null, "eastus", com.azure.core.util.Context.NONE);
    }
}
