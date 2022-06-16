import com.azure.core.util.Context;

/** Samples for Channels ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/ListChannelsByBotService.json
     */
    /**
     * Sample code: List Bots by Resource Group.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void listBotsByResourceGroup(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.channels().listByResourceGroup("OneResourceGroupName", "samplebotname", Context.NONE);
    }
}
