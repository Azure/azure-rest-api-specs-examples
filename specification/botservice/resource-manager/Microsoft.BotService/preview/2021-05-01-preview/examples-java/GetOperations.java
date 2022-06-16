import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/GetOperations.json
     */
    /**
     * Sample code: Get Operations.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void getOperations(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.operations().list(Context.NONE);
    }
}
