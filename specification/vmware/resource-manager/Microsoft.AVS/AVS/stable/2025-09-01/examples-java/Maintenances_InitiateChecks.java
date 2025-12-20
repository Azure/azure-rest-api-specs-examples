
/**
 * Samples for Maintenances InitiateChecks.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Maintenances_InitiateChecks.json
     */
    /**
     * Sample code: Maintenances_InitiateChecks.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void maintenancesInitiateChecks(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.maintenances().initiateChecksWithResponse("group1", "cloud1", "maintenance1",
            com.azure.core.util.Context.NONE);
    }
}
