
/**
 * Samples for BotConnection ListByBotService.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/examples/
     * ListConnectionsByBotService.json
     */
    /**
     * Sample code: List Connection Settings.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void listConnectionSettings(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.botConnections().listByBotService("OneResourceGroupName", "samplebotname",
            com.azure.core.util.Context.NONE);
    }
}
