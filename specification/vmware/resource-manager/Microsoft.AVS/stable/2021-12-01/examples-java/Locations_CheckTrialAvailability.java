import com.azure.core.util.Context;

/** Samples for Locations CheckTrialAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/Locations_CheckTrialAvailability.json
     */
    /**
     * Sample code: Locations_CheckTrialAvailability.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void locationsCheckTrialAvailability(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.locations().checkTrialAvailabilityWithResponse("eastus", Context.NONE);
    }
}
