/** Samples for VerifiedPartners List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/VerifiedPartners_List.json
     */
    /**
     * Sample code: VerifiedPartners_List.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void verifiedPartnersList(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.verifiedPartners().list(null, null, com.azure.core.util.Context.NONE);
    }
}
