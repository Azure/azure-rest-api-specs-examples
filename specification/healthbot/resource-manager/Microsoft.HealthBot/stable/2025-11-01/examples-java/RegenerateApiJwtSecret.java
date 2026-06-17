
/**
 * Samples for Bots RegenerateApiJwtSecret.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/RegenerateApiJwtSecret.json
     */
    /**
     * Sample code: Bot Regenerate API JWT Secret.
     * 
     * @param manager Entry point to HealthbotManager.
     */
    public static void botRegenerateAPIJWTSecret(com.azure.resourcemanager.healthbot.HealthbotManager manager) {
        manager.bots().regenerateApiJwtSecretWithResponse("healthbotClient", "samplebotname",
            com.azure.core.util.Context.NONE);
    }
}
