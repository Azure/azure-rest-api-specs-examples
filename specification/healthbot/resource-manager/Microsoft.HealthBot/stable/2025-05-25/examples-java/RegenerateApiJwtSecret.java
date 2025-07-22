
/**
 * Samples for Bots RegenerateApiJwtSecret.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2025-05-25/examples/RegenerateApiJwtSecret.
     * json
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
