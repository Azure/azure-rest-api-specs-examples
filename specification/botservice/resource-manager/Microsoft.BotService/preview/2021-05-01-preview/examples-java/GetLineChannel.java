
/**
 * Samples for Channels Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/GetLineChannel
     * .json
     */
    /**
     * Sample code: Get Line Channel.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void getLineChannel(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.channels().getWithResponse("OneResourceGroupName", "samplebotname", "LineChannel",
            com.azure.core.util.Context.NONE);
    }
}
