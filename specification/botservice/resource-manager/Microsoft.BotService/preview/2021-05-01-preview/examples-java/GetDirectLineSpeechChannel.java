
/**
 * Samples for Channels Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/
     * GetDirectLineSpeechChannel.json
     */
    /**
     * Sample code: Get DirectLine Speech Channel.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void getDirectLineSpeechChannel(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.channels().getWithResponse("OneResourceGroupName", "samplebotname", "DirectLineSpeechChannel",
            com.azure.core.util.Context.NONE);
    }
}
