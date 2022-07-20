import com.azure.core.util.Context;

/** Samples for Bots List. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2020-12-08/examples/ListBotsBySubscription.json
     */
    /**
     * Sample code: List Bots by Subscription.
     *
     * @param manager Entry point to HealthbotManager.
     */
    public static void listBotsBySubscription(com.azure.resourcemanager.healthbot.HealthbotManager manager) {
        manager.bots().list(Context.NONE);
    }
}
