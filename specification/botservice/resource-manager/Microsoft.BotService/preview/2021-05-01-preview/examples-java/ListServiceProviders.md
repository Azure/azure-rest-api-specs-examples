```java
import com.azure.core.util.Context;

/** Samples for BotConnection ListServiceProviders. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/ListServiceProviders.json
     */
    /**
     * Sample code: List Auth Service Providers.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void listAuthServiceProviders(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.botConnections().listServiceProvidersWithResponse(Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-botservice_1.0.0-beta.4/sdk/botservice/azure-resourcemanager-botservice/README.md) on how to add the SDK to your project and authenticate.
