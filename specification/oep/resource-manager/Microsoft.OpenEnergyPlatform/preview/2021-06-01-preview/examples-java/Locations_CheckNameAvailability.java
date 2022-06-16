import com.azure.core.util.Context;
import com.azure.resourcemanager.oep.models.CheckNameAvailabilityRequest;

/** Samples for Locations CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2021-06-01-preview/examples/Locations_CheckNameAvailability.json
     */
    /**
     * Sample code: Locations_CheckNameAvailability.
     *
     * @param manager Entry point to OepManager.
     */
    public static void locationsCheckNameAvailability(com.azure.resourcemanager.oep.OepManager manager) {
        manager
            .locations()
            .checkNameAvailabilityWithResponse(
                new CheckNameAvailabilityRequest().withName("sample-name").withType("Microsoft.OEP/oepResource"),
                Context.NONE);
    }
}
