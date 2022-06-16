import com.azure.core.util.Context;
import com.azure.resourcemanager.botservice.fluent.models.BotChannelInner;
import com.azure.resourcemanager.botservice.models.AlexaChannel;
import com.azure.resourcemanager.botservice.models.AlexaChannelProperties;
import com.azure.resourcemanager.botservice.models.ChannelName;

/** Samples for Channels Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/UpdateAlexaChannel.json
     */
    /**
     * Sample code: Update Alexa.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void updateAlexa(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager
            .channels()
            .updateWithResponse(
                "OneResourceGroupName",
                "samplebotname",
                ChannelName.ALEXA_CHANNEL,
                new BotChannelInner()
                    .withLocation("global")
                    .withProperties(
                        new AlexaChannel()
                            .withProperties(
                                new AlexaChannelProperties().withAlexaSkillId("XAlexaSkillIdX").withIsEnabled(true))),
                Context.NONE);
    }
}
