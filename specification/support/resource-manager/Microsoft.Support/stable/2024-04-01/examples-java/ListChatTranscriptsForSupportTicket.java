
/**
 * Samples for ChatTranscriptsNoSubscription List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * ListChatTranscriptsForSupportTicket.json
     */
    /**
     * Sample code: List chat transcripts for a no-subscription support ticket.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void
        listChatTranscriptsForANoSubscriptionSupportTicket(com.azure.resourcemanager.support.SupportManager manager) {
        manager.chatTranscriptsNoSubscriptions().list("testticket", com.azure.core.util.Context.NONE);
    }
}
