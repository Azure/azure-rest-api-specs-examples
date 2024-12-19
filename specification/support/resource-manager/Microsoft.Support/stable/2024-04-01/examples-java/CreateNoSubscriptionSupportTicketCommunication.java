
import com.azure.resourcemanager.support.fluent.models.CommunicationDetailsInner;

/**
 * Samples for CommunicationsNoSubscription Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * CreateNoSubscriptionSupportTicketCommunication.json
     */
    /**
     * Sample code: AddCommunicationToNoSubscriptionTicket.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void
        addCommunicationToNoSubscriptionTicket(com.azure.resourcemanager.support.SupportManager manager) {
        manager.communicationsNoSubscriptions().create("testticket", "testcommunication",
            new CommunicationDetailsInner().withSender("user@contoso.com")
                .withSubject("This is a test message from a customer!")
                .withBody("This is a test message from a customer!"),
            com.azure.core.util.Context.NONE);
    }
}
