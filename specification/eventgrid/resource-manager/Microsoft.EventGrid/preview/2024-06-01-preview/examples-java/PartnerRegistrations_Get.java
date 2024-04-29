
/**
 * Samples for PartnerRegistrations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/
     * PartnerRegistrations_Get.json
     */
    /**
     * Sample code: PartnerRegistrations_Get.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerRegistrationsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerRegistrations().getByResourceGroupWithResponse("examplerg", "examplePartnerRegistrationName1",
            com.azure.core.util.Context.NONE);
    }
}
