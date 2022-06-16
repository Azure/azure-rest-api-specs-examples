import com.azure.core.util.Context;

/** Samples for AvailabilityStatuses GetByResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2020-05-01/examples/AvailabilityStatus_GetByResource.json
     */
    /**
     * Sample code: GetCurrentHealthByResource.
     *
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void getCurrentHealthByResource(
        com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager
            .availabilityStatuses()
            .getByResourceWithResponse("resourceUri", null, "recommendedactions", Context.NONE);
    }
}
