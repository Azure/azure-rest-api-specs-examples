import com.azure.core.util.Context;

/** Samples for PartnerDestinations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerDestinations_Delete.json
     */
    /**
     * Sample code: PartnerDestinations_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerDestinationsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerDestinations().delete("examplerg", "examplePartnerDestinationName1", Context.NONE);
    }
}
