
/**
 * Samples for Bots Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/examples/DeleteBot.json
     */
    /**
     * Sample code: Delete Bot.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void deleteBot(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.bots().deleteByResourceGroupWithResponse("OneResourceGroupName", "samplebotname",
            com.azure.core.util.Context.NONE);
    }
}
