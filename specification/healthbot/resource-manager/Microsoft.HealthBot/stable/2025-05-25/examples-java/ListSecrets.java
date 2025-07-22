
/**
 * Samples for Bots ListSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2025-05-25/examples/ListSecrets.json
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
