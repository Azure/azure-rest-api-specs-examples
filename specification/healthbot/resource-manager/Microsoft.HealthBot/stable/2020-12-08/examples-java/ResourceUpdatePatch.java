import com.azure.core.util.Context;
import com.azure.resourcemanager.healthbot.models.HealthBot;
import com.azure.resourcemanager.healthbot.models.Sku;
import com.azure.resourcemanager.healthbot.models.SkuName;

/** Samples for Bots Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2020-12-08/examples/ResourceUpdatePatch.json
     */
    /**
     * Sample code: BotUpdate.
     *
     * @param manager Entry point to HealthbotManager.
     */
    public static void botUpdate(com.azure.resourcemanager.healthbot.HealthbotManager manager) {
        HealthBot resource =
            manager.bots().getByResourceGroupWithResponse("healthbotClient", "samplebotname", Context.NONE).getValue();
        resource.update().withSku(new Sku().withName(SkuName.F0)).apply();
    }
}
