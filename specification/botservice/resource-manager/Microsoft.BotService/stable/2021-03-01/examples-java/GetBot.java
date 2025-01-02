
/**
 * Samples for Bots GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/examples/GetBot.json
     */
    /**
     * Sample code: Get Bot.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void getBot(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.bots().getByResourceGroupWithResponse("OneResourceGroupName", "samplebotname",
            com.azure.core.util.Context.NONE);
    }
}
