import com.azure.core.util.Context;

/** Samples for Bots GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/GetBot.json
     */
    /**
     * Sample code: Get Bot.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void getBot(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.bots().getByResourceGroupWithResponse("OneResourceGroupName", "samplebotname", Context.NONE);
    }
}
