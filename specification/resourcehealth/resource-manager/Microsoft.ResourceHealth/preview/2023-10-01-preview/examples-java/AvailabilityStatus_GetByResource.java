/** Samples for AvailabilityStatuses GetByResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/AvailabilityStatus_GetByResource.json
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
            .getByResourceWithResponse("resourceUri", null, "recommendedactions", com.azure.core.util.Context.NONE);
    }
}
