
import com.azure.resourcemanager.support.models.CheckNameAvailabilityInput;
import com.azure.resourcemanager.support.models.Type;

/**
 * Samples for SupportTickets CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * CheckNameAvailabilityWithSubscription.json
     */
    /**
     * Sample code: Checks whether name is available for a subscription support ticket resource.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void checksWhetherNameIsAvailableForASubscriptionSupportTicketResource(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().checkNameAvailabilityWithResponse(
            new CheckNameAvailabilityInput().withName("sampleName").withType(Type.MICROSOFT_SUPPORT_SUPPORT_TICKETS),
            com.azure.core.util.Context.NONE);
    }
}
