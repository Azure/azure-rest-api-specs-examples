/** Samples for SupportTicketsNoSubscription Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/GetSupportTicketDetails.json
     */
    /**
     * Sample code: Get details of a ticket.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void getDetailsOfATicket(com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTicketsNoSubscriptions().getWithResponse("testticket", com.azure.core.util.Context.NONE);
    }
}
