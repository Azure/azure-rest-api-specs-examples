
/**
 * Samples for Bots List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2025-05-25/examples/ListBotsBySubscription.
     * json
     */
    /**
     * Sample code: List Bots by Subscription.
     * 
     * @param manager Entry point to HealthbotManager.
     */
    public static void listBotsBySubscription(com.azure.resourcemanager.healthbot.HealthbotManager manager) {
        manager.bots().list(com.azure.core.util.Context.NONE);
    }
}
