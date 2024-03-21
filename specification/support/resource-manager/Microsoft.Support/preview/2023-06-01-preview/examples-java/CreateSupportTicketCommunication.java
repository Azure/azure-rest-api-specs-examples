
/**
 * Samples for Communications Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * CreateSupportTicketCommunication.json
     */
    /**
     * Sample code: AddCommunicationToSubscriptionTicket.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void addCommunicationToSubscriptionTicket(com.azure.resourcemanager.support.SupportManager manager) {
        manager.communications().define("testcommunication").withExistingSupportTicket("testticket")
            .withSender("user@contoso.com").withSubject("This is a test message from a customer!")
            .withBody("This is a test message from a customer!").create();
    }
}
