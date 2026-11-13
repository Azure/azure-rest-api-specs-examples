
/**
 * Samples for ManagedComputeCapacities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ListManagedComputeCapacitiesWithUnsupportedLocation.json
     */
    /**
     * Sample code: List Managed Compute Capacities With Unsupported DataZone Location.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listManagedComputeCapacitiesWithUnsupportedDataZoneLocation(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.managedComputeCapacities().list("MaaP", "H100_80GB", null, "japaneast",
            com.azure.core.util.Context.NONE);
    }
}
