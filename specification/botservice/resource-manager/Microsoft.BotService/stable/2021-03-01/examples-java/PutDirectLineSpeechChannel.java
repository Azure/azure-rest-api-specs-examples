
import com.azure.resourcemanager.botservice.fluent.models.BotChannelInner;
import com.azure.resourcemanager.botservice.models.ChannelName;
import com.azure.resourcemanager.botservice.models.DirectLineSpeechChannel;
import com.azure.resourcemanager.botservice.models.DirectLineSpeechChannelProperties;

/**
 * Samples for Channels Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/examples/
     * PutDirectLineSpeechChannel.json
     */
    /**
     * Sample code: Create DirectLine Speech Channel.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void createDirectLineSpeechChannel(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.channels().createWithResponse("OneResourceGroupName", "samplebotname",
            ChannelName.DIRECT_LINE_SPEECH_CHANNEL,
            new BotChannelInner().withLocation("global")
                .withProperties(new DirectLineSpeechChannel().withProperties(
                    new DirectLineSpeechChannelProperties().withCognitiveServiceRegion("XcognitiveServiceRegionX")
                        .withCognitiveServiceSubscriptionKey("fakeTokenPlaceholder").withIsEnabled(true))),
            com.azure.core.util.Context.NONE);
    }
}
