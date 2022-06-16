import com.azure.core.util.Context;
import com.azure.resourcemanager.botservice.fluent.models.BotChannelInner;
import com.azure.resourcemanager.botservice.models.ChannelName;
import com.azure.resourcemanager.botservice.models.EmailChannel;
import com.azure.resourcemanager.botservice.models.EmailChannelProperties;

/** Samples for Channels Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/PutChannel.json
     */
    /**
     * Sample code: Create Bot.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void createBot(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager
            .channels()
            .createWithResponse(
                "OneResourceGroupName",
                "samplebotname",
                ChannelName.EMAIL_CHANNEL,
                new BotChannelInner()
                    .withLocation("global")
                    .withProperties(
                        new EmailChannel()
                            .withProperties(
                                new EmailChannelProperties()
                                    .withEmailAddress("a@b.com")
                                    .withPassword("pwd")
                                    .withIsEnabled(true))),
                Context.NONE);
    }
}
