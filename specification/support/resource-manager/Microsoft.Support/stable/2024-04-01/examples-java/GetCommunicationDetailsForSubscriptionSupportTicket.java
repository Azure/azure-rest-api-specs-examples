
/**
 * Samples for Communications Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * GetCommunicationDetailsForSubscriptionSupportTicket.json
     */
    /**
     * Sample code: Get communication details for a subscription support ticket.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void
        getCommunicationDetailsForASubscriptionSupportTicket(com.azure.resourcemanager.support.SupportManager manager) {
        manager.communications().getWithResponse("testticket", "testmessage", com.azure.core.util.Context.NONE);
    }
}
