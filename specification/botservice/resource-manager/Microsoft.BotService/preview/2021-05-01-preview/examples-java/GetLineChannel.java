import com.azure.core.util.Context;

/** Samples for Channels Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/GetLineChannel.json
     */
    /**
     * Sample code: Get Line Bot.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void getLineBot(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.channels().getWithResponse("OneResourceGroupName", "samplebotname", "LineChannel", Context.NONE);
    }
}
