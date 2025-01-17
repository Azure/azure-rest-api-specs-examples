
/**
 * Samples for Channels ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/examples/
     * ListChannelsByBotService.json
     */
    /**
     * Sample code: List Channels by Resource Group.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void listChannelsByResourceGroup(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.channels().listByResourceGroup("OneResourceGroupName", "samplebotname",
            com.azure.core.util.Context.NONE);
    }
}
