/** Samples for Communications List. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/ListWebCommunicationsForSubscriptionSupportTicketCreatedOnOrAfter.json
     */
    /**
     * Sample code: List web communication created on or after a specific date for a subscription support ticket.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void listWebCommunicationCreatedOnOrAfterASpecificDateForASubscriptionSupportTicket(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager
            .communications()
            .list(
                "testticket",
                null,
                "communicationType eq 'web' and createdDate ge 2020-03-10T22:08:51Z",
                com.azure.core.util.Context.NONE);
    }
}
