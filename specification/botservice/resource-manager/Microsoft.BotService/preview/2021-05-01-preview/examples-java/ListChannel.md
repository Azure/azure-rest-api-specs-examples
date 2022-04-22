Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-botservice_1.0.0-beta.4/sdk/botservice/azure-resourcemanager-botservice/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.botservice.models.ChannelName;

/** Samples for Channels ListWithKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/ListChannel.json
     */
    /**
     * Sample code: List Channel.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void listChannel(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager
            .channels()
            .listWithKeysWithResponse("OneResourceGroupName", "samplebotname", ChannelName.EMAIL_CHANNEL, Context.NONE);
    }
}
```
