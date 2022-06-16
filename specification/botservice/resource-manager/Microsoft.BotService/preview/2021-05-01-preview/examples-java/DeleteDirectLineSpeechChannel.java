import com.azure.core.util.Context;

/** Samples for Channels Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/DeleteDirectLineSpeechChannel.json
     */
    /**
     * Sample code: Delete DirectLine Speech Bot.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void deleteDirectLineSpeechBot(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager
            .channels()
            .deleteWithResponse("OneResourceGroupName", "samplebotname", "DirectLineSpeechChannel", Context.NONE);
    }
}
