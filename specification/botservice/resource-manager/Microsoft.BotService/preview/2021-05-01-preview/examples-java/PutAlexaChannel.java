
import com.azure.resourcemanager.botservice.fluent.models.BotChannelInner;
import com.azure.resourcemanager.botservice.models.AlexaChannel;
import com.azure.resourcemanager.botservice.models.AlexaChannelProperties;
import com.azure.resourcemanager.botservice.models.ChannelName;

/**
 * Samples for Channels Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/
     * PutAlexaChannel.json
     */
    /**
     * Sample code: Create Alexa Channel.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void createAlexaChannel(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.channels().createWithResponse("OneResourceGroupName", "samplebotname", ChannelName.ALEXA_CHANNEL,
            new BotChannelInner().withLocation("global")
                .withProperties(new AlexaChannel().withProperties(
                    new AlexaChannelProperties().withAlexaSkillId("XAlexaSkillIdX").withIsEnabled(true))),
            com.azure.core.util.Context.NONE);
    }
}
