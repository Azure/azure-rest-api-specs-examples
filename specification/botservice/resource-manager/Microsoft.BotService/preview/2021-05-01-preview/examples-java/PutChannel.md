Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-botservice_1.0.0-beta.2/sdk/botservice/azure-resourcemanager-botservice/README.md) on how to add the SDK to your project and authenticate.

```java
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
            .define(ChannelName.EMAIL_CHANNEL)
            .withRegion("global")
            .withExistingBotService("OneResourceGroupName", "samplebotname")
            .withProperties(
                new EmailChannel()
                    .withProperties(
                        new EmailChannelProperties()
                            .withEmailAddress("a@b.com")
                            .withPassword("pwd")
                            .withIsEnabled(true)))
            .create();
    }
}
```
