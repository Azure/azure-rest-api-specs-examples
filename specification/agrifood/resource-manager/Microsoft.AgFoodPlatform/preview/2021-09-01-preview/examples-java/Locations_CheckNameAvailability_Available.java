import com.azure.core.util.Context;
import com.azure.resourcemanager.agrifood.models.CheckNameAvailabilityRequest;

/** Samples for Locations CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/Locations_CheckNameAvailability_Available.json
     */
    /**
     * Sample code: Locations_CheckNameAvailability_Available.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void locationsCheckNameAvailabilityAvailable(
        com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager
            .locations()
            .checkNameAvailabilityWithResponse(
                new CheckNameAvailabilityRequest()
                    .withName("newaccountname")
                    .withType("Microsoft.AgFoodPlatform/farmBeats"),
                Context.NONE);
    }
}
