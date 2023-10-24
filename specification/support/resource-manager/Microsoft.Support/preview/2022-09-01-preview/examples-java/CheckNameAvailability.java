import com.azure.resourcemanager.support.models.CheckNameAvailabilityInput;
import com.azure.resourcemanager.support.models.Type;

/** Samples for SupportTicketsNoSubscription CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/CheckNameAvailability.json
     */
    /**
     * Sample code: Checks whether name is available for SupportTicket resource.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void checksWhetherNameIsAvailableForSupportTicketResource(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager
            .supportTicketsNoSubscriptions()
            .checkNameAvailabilityWithResponse(
                new CheckNameAvailabilityInput()
                    .withName("sampleName")
                    .withType(Type.MICROSOFT_SUPPORT_SUPPORT_TICKETS),
                com.azure.core.util.Context.NONE);
    }
}
