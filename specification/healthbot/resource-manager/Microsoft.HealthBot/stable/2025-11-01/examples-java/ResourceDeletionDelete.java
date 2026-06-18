
/**
 * Samples for Bots Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/ResourceDeletionDelete.json
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
