
import com.azure.resourcemanager.appservice.fluent.models.SitePatchResourceInner;

/**
 * Samples for WebApps UpdateSlot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/UpdateWebAppSlot.json
     */
    /**
     * Sample code: Update Web App Slot.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateWebAppSlot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps()
            .updateSlotWithResponse("testrg123", "sitef6141", "staging", new SitePatchResourceInner().withServerFarmId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/DefaultAsp"),
                com.azure.core.util.Context.NONE);
    }
}
