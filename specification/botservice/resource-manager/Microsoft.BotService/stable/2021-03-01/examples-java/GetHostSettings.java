
/**
 * Samples for HostSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/examples/GetHostSettings.json
     */
    /**
     * Sample code: Get Bot Host Settings.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void getBotHostSettings(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.hostSettings().getWithResponse(com.azure.core.util.Context.NONE);
    }
}
