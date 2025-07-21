
/**
 * Samples for Bots GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2025-05-25/examples/ResourceInfoGet.json
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
