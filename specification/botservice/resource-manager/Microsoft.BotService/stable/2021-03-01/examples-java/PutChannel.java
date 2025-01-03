
import com.azure.resourcemanager.botservice.fluent.models.BotChannelInner;
import com.azure.resourcemanager.botservice.models.ChannelName;
import com.azure.resourcemanager.botservice.models.EmailChannel;
import com.azure.resourcemanager.botservice.models.EmailChannelProperties;

/**
 * Samples for Channels Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/examples/PutChannel.json
     */
    /**
     * Sample code: Create Channel.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void createChannel(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.channels().createWithResponse("OneResourceGroupName", "samplebotname", ChannelName.EMAIL_CHANNEL,
            new BotChannelInner().withLocation("global")
                .withProperties(new EmailChannel().withProperties(new EmailChannelProperties()
                    .withEmailAddress("a@b.com").withPassword("fakeTokenPlaceholder").withIsEnabled(true))),
            com.azure.core.util.Context.NONE);
    }
}
