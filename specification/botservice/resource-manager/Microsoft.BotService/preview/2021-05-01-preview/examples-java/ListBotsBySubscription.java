import com.azure.core.util.Context;

/** Samples for Bots List. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/ListBotsBySubscription.json
     */
    /**
     * Sample code: List Bots by Subscription.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void listBotsBySubscription(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.bots().list(Context.NONE);
    }
}
