
/**
 * Samples for CommunicationsNoSubscription Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * GetCommunicationDetailsForSupportTicket.json
     */
    /**
     * Sample code: Get communication details for a no-subscription support ticket.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void getCommunicationDetailsForANoSubscriptionSupportTicket(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.communicationsNoSubscriptions().getWithResponse("testticket", "testmessage",
            com.azure.core.util.Context.NONE);
    }
}
