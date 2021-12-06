Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-botservice_1.0.0-beta.2/sdk/botservice/azure-resourcemanager-botservice/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.botservice.models.BotChannel;
import com.azure.resourcemanager.botservice.models.LineChannel;
import com.azure.resourcemanager.botservice.models.LineChannelProperties;
import com.azure.resourcemanager.botservice.models.LineRegistration;
import java.util.Arrays;

/** Samples for Channels Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/UpdateLineChannel.json
     */
    /**
     * Sample code: Update Line.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void updateLine(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        BotChannel resource =
            manager
                .channels()
                .getWithResponse("OneResourceGroupName", "samplebotname", "LineChannel", Context.NONE)
                .getValue();
        resource
            .update()
            .withProperties(
                new LineChannel()
                    .withProperties(
                        new LineChannelProperties()
                            .withLineRegistrations(
                                Arrays
                                    .asList(
                                        new LineRegistration()
                                            .withChannelSecret("channelSecret")
                                            .withChannelAccessToken("channelAccessToken")))))
            .apply();
    }
}
```
