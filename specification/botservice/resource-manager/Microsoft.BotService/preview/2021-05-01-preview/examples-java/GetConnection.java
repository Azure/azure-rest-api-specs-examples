
/**
 * Samples for BotConnection ListWithSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/GetConnection.
     * json
     */
    /**
     * Sample code: List Connection Setting With Secrets.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void
        listConnectionSettingWithSecrets(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.botConnections().listWithSecretsWithResponse("OneResourceGroupName", "samplebotname",
            "sampleConnection", com.azure.core.util.Context.NONE);
    }
}
