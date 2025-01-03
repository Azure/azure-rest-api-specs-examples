
import com.azure.resourcemanager.botservice.models.ChannelName;

/**
 * Samples for Channels ListWithKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/examples/ListChannel.json
     */
    /**
     * Sample code: List Channel.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void listChannel(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.channels().listWithKeysWithResponse("OneResourceGroupName", "samplebotname", ChannelName.EMAIL_CHANNEL,
            com.azure.core.util.Context.NONE);
    }
}
