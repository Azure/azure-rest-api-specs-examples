Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-botservice_1.0.0-beta.2/sdk/botservice/azure-resourcemanager-botservice/README.md) on how to add the SDK to your project and authenticate.

```java
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
            .define(ChannelName.ALEXA_CHANNEL)
            .withRegion("global")
            .withExistingBotService("OneResourceGroupName", "samplebotname")
            .withProperties(
                new AlexaChannel()
                    .withProperties(
                        new AlexaChannelProperties().withAlexaSkillId("XAlexaSkillIdX").withIsEnabled(true)))
            .create();
    }
}
```
