
/**
 * Samples for Bots ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/ListBotsByResourceGroup.json
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
