
/**
 * Samples for AvailabilityStatuses ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/
     * AvailabilityStatuses_ListByResourceGroup.json
     */
    /**
     * Sample code: ListByResourceGroup.
     * 
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void listByResourceGroup(com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager.availabilityStatuses().listByResourceGroup("resourceGroupName", null, "recommendedactions",
            com.azure.core.util.Context.NONE);
    }
}
