
/**
 * Samples for BotConnection ListServiceProviders.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/examples/ListServiceProviders.
     * json
     */
    /**
     * Sample code: List Auth Service Providers.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void listAuthServiceProviders(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.botConnections().listServiceProvidersWithResponse(com.azure.core.util.Context.NONE);
    }
}
