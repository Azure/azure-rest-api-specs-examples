
import com.azure.resourcemanager.appservice.models.CsmSlotEntity;

/**
 * Samples for WebApps ApplySlotConfigToProduction.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/ApplySlotConfig.json
     */
    /**
     * Sample code: Apply web app slot config.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void applyWebAppSlotConfig(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().applySlotConfigToProductionWithResponse("testrg123",
            "sitef6141", new CsmSlotEntity().withTargetSlot("staging").withPreserveVnet(true),
            com.azure.core.util.Context.NONE);
    }
}
