
/**
 * Samples for ChatTranscripts List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * ListChatTranscriptsForSubscriptionSupportTicket.json
     */
    /**
     * Sample code: List chat transcripts for a subscription support ticket.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void
        listChatTranscriptsForASubscriptionSupportTicket(com.azure.resourcemanager.support.SupportManager manager) {
        manager.chatTranscripts().list("testticket", com.azure.core.util.Context.NONE);
    }
}
