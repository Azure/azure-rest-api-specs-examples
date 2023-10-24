import com.azure.resourcemanager.support.models.CheckNameAvailabilityInput;
import com.azure.resourcemanager.support.models.Type;

/** Samples for Communications CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/CheckNameAvailabilityForSupportTicketCommunication.json
     */
    /**
     * Sample code: Checks whether name is available for Communication resource.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void checksWhetherNameIsAvailableForCommunicationResource(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager
            .communications()
            .checkNameAvailabilityWithResponse(
                "testticket",
                new CheckNameAvailabilityInput().withName("sampleName").withType(Type.MICROSOFT_SUPPORT_COMMUNICATIONS),
                com.azure.core.util.Context.NONE);
    }
}
