
/**
 * Samples for Channels Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/examples/GetAlexaChannel.json
     */
    /**
     * Sample code: Get Alexa Channel.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void getAlexaChannel(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.channels().getWithResponse("OneResourceGroupName", "samplebotname", "AlexaChannel",
            com.azure.core.util.Context.NONE);
    }
}
