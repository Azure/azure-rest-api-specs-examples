```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.botservice.fluent.models.BotChannelInner;
import com.azure.resourcemanager.botservice.models.ChannelName;
import com.azure.resourcemanager.botservice.models.LineChannel;
import com.azure.resourcemanager.botservice.models.LineChannelProperties;
import com.azure.resourcemanager.botservice.models.LineRegistration;
import java.util.Arrays;

/** Samples for Channels Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/PutLineChannel.json
     */
    /**
     * Sample code: Create Line Bot.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void createLineBot(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager
            .channels()
            .createWithResponse(
                "OneResourceGroupName",
                "samplebotname",
                ChannelName.LINE_CHANNEL,
                new BotChannelInner()
                    .withLocation("global")
                    .withProperties(
                        new LineChannel()
                            .withProperties(
                                new LineChannelProperties()
                                    .withLineRegistrations(
                                        Arrays
                                            .asList(
                                                new LineRegistration()
                                                    .withChannelSecret("channelSecret")
                                                    .withChannelAccessToken("channelAccessToken"))))),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-botservice_1.0.0-beta.4/sdk/botservice/azure-resourcemanager-botservice/README.md) on how to add the SDK to your project and authenticate.
