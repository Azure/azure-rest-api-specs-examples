import com.azure.core.util.Context;

/** Samples for PartnerRegistrations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/PartnerRegistrations_Delete.json
     */
    /**
     * Sample code: PartnerRegistrations_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerRegistrationsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerRegistrations().deleteWithResponse("examplerg", "examplePartnerRegistrationName1", Context.NONE);
    }
}
