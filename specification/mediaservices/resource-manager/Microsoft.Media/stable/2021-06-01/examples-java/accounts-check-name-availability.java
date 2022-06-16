import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.CheckNameAvailabilityInput;

/** Samples for Locations CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/accounts-check-name-availability.json
     */
    /**
     * Sample code: Check Name Availability.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void checkNameAvailability(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .locations()
            .checkNameAvailabilityWithResponse(
                "japanwest",
                new CheckNameAvailabilityInput().withName("contosotv").withType("videoAnalyzers"),
                Context.NONE);
    }
}
