
/**
 * Samples for ChatTranscriptsNoSubscription Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * GetchatTranscriptDetailsForSupportTicket.json
     */
    /**
     * Sample code: Get chat transcript details for a subscription support ticket.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void getChatTranscriptDetailsForASubscriptionSupportTicket(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.chatTranscriptsNoSubscriptions().getWithResponse("testticket", "b371192a-b094-4a71-b093-7246029b0a54",
            com.azure.core.util.Context.NONE);
    }
}
