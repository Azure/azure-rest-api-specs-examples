
/**
 * Samples for VerifiedPartners Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * VerifiedPartners_Get.json
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
