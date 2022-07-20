import com.azure.resourcemanager.healthbot.models.Sku;
import com.azure.resourcemanager.healthbot.models.SkuName;

/** Samples for Bots Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2020-12-08/examples/ResourceCreationPut.json
     */
    /**
     * Sample code: BotCreate.
     *
     * @param manager Entry point to HealthbotManager.
     */
    public static void botCreate(com.azure.resourcemanager.healthbot.HealthbotManager manager) {
        manager
            .bots()
            .define("samplebotname")
            .withRegion("East US")
            .withExistingResourceGroup("healthbotClient")
            .withSku(new Sku().withName(SkuName.F0))
            .create();
    }
}
