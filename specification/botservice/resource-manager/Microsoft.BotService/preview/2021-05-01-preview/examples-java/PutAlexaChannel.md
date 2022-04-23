Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-botservice_1.0.0-beta.4/sdk/botservice/azure-resourcemanager-botservice/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.botservice.fluent.models.BotChannelInner;
import com.azure.resourcemanager.botservice.models.AlexaChannel;
import com.azure.resourcemanager.botservice.models.AlexaChannelProperties;
import com.azure.resourcemanager.botservice.models.ChannelName;

/** Samples for Channels Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/PutAlexaChannel.json
     */
    /**
     * Sample code: Create Alexa Bot.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void createAlexaBot(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager
            .channels()
            .createWithResponse(
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
```
