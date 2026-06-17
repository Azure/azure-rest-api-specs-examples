
/**
 * Samples for Bots ListSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/ListSecrets.json
     */
    /**
     * Sample code: Bot List Secrets.
     * 
     * @param manager Entry point to HealthbotManager.
     */
    public static void botListSecrets(com.azure.resourcemanager.healthbot.HealthbotManager manager) {
        manager.bots().listSecretsWithResponse("healthbotClient", "samplebotname", com.azure.core.util.Context.NONE);
    }
}
