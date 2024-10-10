
/**
 * Samples for Channels Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/DeleteChannel.
     * json
     */
    /**
     * Sample code: Delete Channel.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void deleteChannel(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.channels().deleteWithResponse("OneResourceGroupName", "samplebotname", "EmailChannel",
            com.azure.core.util.Context.NONE);
    }
}
