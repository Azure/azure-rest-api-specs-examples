
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/examples/GetOperations.json
     */
    /**
     * Sample code: Get Operations.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void getOperations(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
