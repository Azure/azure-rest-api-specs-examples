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
