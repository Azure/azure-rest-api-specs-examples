import com.azure.core.util.Context;

/** Samples for Channels Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/GetAlexaChannel.json
     */
    /**
     * Sample code: Get Alexa Bot.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void getAlexaBot(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.channels().getWithResponse("OneResourceGroupName", "samplebotname", "AlexaChannel", Context.NONE);
    }
}
