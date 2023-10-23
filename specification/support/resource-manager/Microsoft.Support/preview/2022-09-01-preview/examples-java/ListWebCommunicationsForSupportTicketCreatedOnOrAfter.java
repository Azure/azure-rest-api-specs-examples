/** Samples for SupportTicketCommunicationsNoSubscription List. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/ListWebCommunicationsForSupportTicketCreatedOnOrAfter.json
     */
    /**
     * Sample code: List web communication created on or after a specific date for a no-subscription support ticket.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void listWebCommunicationCreatedOnOrAfterASpecificDateForANoSubscriptionSupportTicket(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager
            .supportTicketCommunicationsNoSubscriptions()
            .list(
                "testticket",
                null,
                "communicationType eq 'web' and createdDate ge 2020-03-10T22:08:51Z",
                com.azure.core.util.Context.NONE);
    }
}
