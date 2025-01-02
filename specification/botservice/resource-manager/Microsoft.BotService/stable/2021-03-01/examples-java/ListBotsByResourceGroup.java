
/**
 * Samples for Bots ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/examples/ListBotsByResourceGroup
     * .json
     */
    /**
     * Sample code: List Bots by Resource Group.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void listBotsByResourceGroup(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.bots().listByResourceGroup("OneResourceGroupName", com.azure.core.util.Context.NONE);
    }
}
