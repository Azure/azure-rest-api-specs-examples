import com.azure.core.util.Context;

/** Samples for PrivateLinkResources ListByBotResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/ListPrivateLinkResources.json
     */
    /**
     * Sample code: List Private Link Resources.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void listPrivateLinkResources(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.privateLinkResources().listByBotResourceWithResponse("res6977", "sto2527", Context.NONE);
    }
}
