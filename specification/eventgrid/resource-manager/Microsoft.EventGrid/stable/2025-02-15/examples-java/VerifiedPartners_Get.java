
/**
 * Samples for VerifiedPartners Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/VerifiedPartners_Get.json
     */
    /**
     * Sample code: VerifiedPartners_Get.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void verifiedPartnersGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.verifiedPartners().getWithResponse("Contoso.Finance", com.azure.core.util.Context.NONE);
    }
}
