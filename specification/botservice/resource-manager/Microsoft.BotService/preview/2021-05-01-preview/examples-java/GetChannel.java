
/**
 * Samples for Channels Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/GetChannel.
     * json
     */
    /**
     * Sample code: Get Channel.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void getChannel(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.channels().getWithResponse("OneResourceGroupName", "samplebotname", "EmailChannel",
            com.azure.core.util.Context.NONE);
    }
}
