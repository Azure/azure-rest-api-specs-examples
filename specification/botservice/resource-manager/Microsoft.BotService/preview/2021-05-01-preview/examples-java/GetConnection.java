import com.azure.core.util.Context;

/** Samples for BotConnection Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/GetConnection.json
     */
    /**
     * Sample code: Update Connection Setting.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void updateConnectionSetting(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager
            .botConnections()
            .getWithResponse("OneResourceGroupName", "samplebotname", "sampleConnection", Context.NONE);
    }
}
