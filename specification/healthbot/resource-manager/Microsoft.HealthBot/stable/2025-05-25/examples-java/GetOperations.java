
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2025-05-25/examples/GetOperations.json
     */
    /**
     * Sample code: Get Operations.
     * 
     * @param manager Entry point to HealthbotManager.
     */
    public static void getOperations(com.azure.resourcemanager.healthbot.HealthbotManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
