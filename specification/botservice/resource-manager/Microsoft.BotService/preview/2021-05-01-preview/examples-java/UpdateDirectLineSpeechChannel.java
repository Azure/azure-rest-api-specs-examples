
import com.azure.resourcemanager.botservice.fluent.models.BotChannelInner;
import com.azure.resourcemanager.botservice.models.ChannelName;
import com.azure.resourcemanager.botservice.models.DirectLineSpeechChannel;
import com.azure.resourcemanager.botservice.models.DirectLineSpeechChannelProperties;

/**
 * Samples for Channels Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/
     * UpdateDirectLineSpeechChannel.json
     */
    /**
     * Sample code: Update DirectLine Speech Channel.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void updateDirectLineSpeechChannel(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.channels().updateWithResponse("OneResourceGroupName", "samplebotname",
            ChannelName.DIRECT_LINE_SPEECH_CHANNEL,
            new BotChannelInner().withLocation("global")
                .withProperties(new DirectLineSpeechChannel().withProperties(
                    new DirectLineSpeechChannelProperties().withCognitiveServiceRegion("XcognitiveServiceRegionX")
                        .withCognitiveServiceSubscriptionKey("fakeTokenPlaceholder").withIsEnabled(true))),
            com.azure.core.util.Context.NONE);
    }
}
