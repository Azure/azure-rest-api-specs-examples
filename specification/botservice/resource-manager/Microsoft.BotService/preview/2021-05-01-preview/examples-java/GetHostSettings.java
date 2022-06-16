import com.azure.core.util.Context;

/** Samples for HostSettings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/GetHostSettings.json
     */
    /**
     * Sample code: Get Bot Host Settings.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void getBotHostSettings(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.hostSettings().getWithResponse(Context.NONE);
    }
}
