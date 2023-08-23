/** Samples for Locations CheckQuotaAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/Locations_CheckQuotaAvailability.json
     */
    /**
     * Sample code: Locations_CheckQuotaAvailability.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void locationsCheckQuotaAvailability(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.locations().checkQuotaAvailabilityWithResponse("eastus", com.azure.core.util.Context.NONE);
    }
}
