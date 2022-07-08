import com.azure.core.util.Context;

/** Samples for VerifiedPartners List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/VerifiedPartners_List.json
     */
    /**
     * Sample code: VerifiedPartners_List.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void verifiedPartnersList(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.verifiedPartners().list(null, null, Context.NONE);
    }
}
