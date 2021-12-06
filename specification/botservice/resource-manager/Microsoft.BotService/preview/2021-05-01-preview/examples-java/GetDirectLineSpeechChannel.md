Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-botservice_1.0.0-beta.2/sdk/botservice/azure-resourcemanager-botservice/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Channels Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/GetDirectLineSpeechChannel.json
     */
    /**
     * Sample code: Get DirectLine Speech Bot.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void getDirectLineSpeechBot(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager
            .channels()
            .getWithResponse("OneResourceGroupName", "samplebotname", "DirectLineSpeechChannel", Context.NONE);
    }
}
```
