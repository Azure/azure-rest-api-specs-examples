
/**
 * Samples for Bots Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2025-05-25/examples/ResourceDeletionDelete.
     * json
     */
    /**
     * Sample code: BotDelete.
     * 
     * @param manager Entry point to HealthbotManager.
     */
    public static void botDelete(com.azure.resourcemanager.healthbot.HealthbotManager manager) {
        manager.bots().delete("healthbotClient", "samplebotname", com.azure.core.util.Context.NONE);
    }
}
