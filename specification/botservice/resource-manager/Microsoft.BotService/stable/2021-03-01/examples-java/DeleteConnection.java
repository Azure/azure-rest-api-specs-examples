
/**
 * Samples for BotConnection Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/examples/DeleteConnection.json
     */
    /**
     * Sample code: Delete Connection Setting.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void deleteConnectionSetting(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.botConnections().deleteWithResponse("OneResourceGroupName", "samplebotname", "sampleConnection",
            com.azure.core.util.Context.NONE);
    }
}
