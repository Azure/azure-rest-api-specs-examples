import com.azure.core.util.Context;
import com.azure.resourcemanager.botservice.fluent.models.BotChannelInner;
import com.azure.resourcemanager.botservice.models.ChannelName;
import com.azure.resourcemanager.botservice.models.DirectLineSpeechChannel;
import com.azure.resourcemanager.botservice.models.DirectLineSpeechChannelProperties;

/** Samples for Channels Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/PutDirectLineSpeechChannel.json
     */
    /**
     * Sample code: Create DirectLine Speech Bot.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void createDirectLineSpeechBot(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager
            .channels()
            .createWithResponse(
                "OneResourceGroupName",
                "samplebotname",
                ChannelName.DIRECT_LINE_SPEECH_CHANNEL,
                new BotChannelInner()
                    .withLocation("global")
                    .withProperties(
                        new DirectLineSpeechChannel()
                            .withProperties(
                                new DirectLineSpeechChannelProperties()
                                    .withCognitiveServiceRegion("XcognitiveServiceRegionX")
                                    .withCognitiveServiceSubscriptionKey("XcognitiveServiceSubscriptionKeyX")
                                    .withIsEnabled(true))),
                Context.NONE);
    }
}
