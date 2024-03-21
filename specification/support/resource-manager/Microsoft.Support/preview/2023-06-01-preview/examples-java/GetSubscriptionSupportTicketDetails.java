
/**
 * Samples for SupportTickets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * GetSubscriptionSupportTicketDetails.json
     */
    /**
     * Sample code: Get details of a subscription ticket.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void getDetailsOfASubscriptionTicket(com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().getWithResponse("testticket", com.azure.core.util.Context.NONE);
    }
}
