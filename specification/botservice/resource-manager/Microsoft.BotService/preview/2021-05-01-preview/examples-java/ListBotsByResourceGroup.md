Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-botservice_1.0.0-beta.4/sdk/botservice/azure-resourcemanager-botservice/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Bots ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/ListBotsByResourceGroup.json
     */
    /**
     * Sample code: List Bots by Resource Group.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void listBotsByResourceGroup(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.bots().listByResourceGroup("OneResourceGroupName", Context.NONE);
    }
}
```
