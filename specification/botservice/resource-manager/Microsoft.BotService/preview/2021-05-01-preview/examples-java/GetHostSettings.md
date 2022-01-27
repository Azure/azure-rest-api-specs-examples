Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-botservice_1.0.0-beta.3/sdk/botservice/azure-resourcemanager-botservice/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for HostSettings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/GetHostSettings.json
     */
    /**
     * Sample code: Get Bot Host Settings.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void getBotHostSettings(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.hostSettings().getWithResponse(Context.NONE);
    }
}
```
