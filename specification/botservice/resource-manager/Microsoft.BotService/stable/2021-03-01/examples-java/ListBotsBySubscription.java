
/**
 * Samples for Bots List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/examples/ListBotsBySubscription.
     * json
     */
    /**
     * Sample code: List Bots by Subscription.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void listBotsBySubscription(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.bots().list(com.azure.core.util.Context.NONE);
    }
}
