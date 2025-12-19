
/**
 * Samples for Locations CheckTrialAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Locations_CheckTrialAvailability.json
     */
    /**
     * Sample code: Locations_CheckTrialAvailability.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void locationsCheckTrialAvailability(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.locations().checkTrialAvailabilityWithResponse("eastus", null, com.azure.core.util.Context.NONE);
    }
}
