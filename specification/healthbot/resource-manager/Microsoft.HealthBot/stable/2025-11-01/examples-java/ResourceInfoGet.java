
/**
 * Samples for Bots GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/ResourceInfoGet.json
     */
    /**
     * Sample code: ResourceInfoGet.
     * 
     * @param manager Entry point to HealthbotManager.
     */
    public static void resourceInfoGet(com.azure.resourcemanager.healthbot.HealthbotManager manager) {
        manager.bots().getByResourceGroupWithResponse("healthbotClient", "samplebotname",
            com.azure.core.util.Context.NONE);
    }
}
