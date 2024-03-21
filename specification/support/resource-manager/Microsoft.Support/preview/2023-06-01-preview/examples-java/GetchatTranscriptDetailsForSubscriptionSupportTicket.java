
/**
 * Samples for ChatTranscripts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * GetchatTranscriptDetailsForSubscriptionSupportTicket.json
     */
    /**
     * Sample code: Get chat transcript details for a subscription support ticket.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void getChatTranscriptDetailsForASubscriptionSupportTicket(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.chatTranscripts().getWithResponse("testticket", "69586795-45e9-45b5-bd9e-c9bb237d3e44",
            com.azure.core.util.Context.NONE);
    }
}
