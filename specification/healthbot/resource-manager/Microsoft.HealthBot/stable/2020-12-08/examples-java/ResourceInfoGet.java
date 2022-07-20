import com.azure.core.util.Context;

/** Samples for Bots GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2020-12-08/examples/ResourceInfoGet.json
     */
    /**
     * Sample code: ResourceInfoGet.
     *
     * @param manager Entry point to HealthbotManager.
     */
    public static void resourceInfoGet(com.azure.resourcemanager.healthbot.HealthbotManager manager) {
        manager.bots().getByResourceGroupWithResponse("healthbotClient", "samplebotname", Context.NONE);
    }
}
