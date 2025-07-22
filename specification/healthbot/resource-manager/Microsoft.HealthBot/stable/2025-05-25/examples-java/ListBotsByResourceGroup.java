
/**
 * Samples for Bots ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2025-05-25/examples/ListBotsByResourceGroup.
     * json
     */
    /**
     * Sample code: List Bots by Resource Group.
     * 
     * @param manager Entry point to HealthbotManager.
     */
    public static void listBotsByResourceGroup(com.azure.resourcemanager.healthbot.HealthbotManager manager) {
        manager.bots().listByResourceGroup("OneResourceGroupName", com.azure.core.util.Context.NONE);
    }
}
