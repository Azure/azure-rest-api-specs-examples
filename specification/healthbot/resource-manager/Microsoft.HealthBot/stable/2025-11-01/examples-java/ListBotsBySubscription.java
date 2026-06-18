
/**
 * Samples for Bots List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/ListBotsBySubscription.json
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
