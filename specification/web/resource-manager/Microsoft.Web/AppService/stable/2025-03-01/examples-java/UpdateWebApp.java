
import com.azure.resourcemanager.appservice.fluent.models.SitePatchResourceInner;

/**
 * Samples for WebApps Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/UpdateWebApp.json
     */
    /**
     * Sample code: Update web app.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateWebApp(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().update("testrg123", "sitef6141",
            new SitePatchResourceInner().withServerFarmId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/DefaultAsp"),
            com.azure.core.util.Context.NONE);
    }
}
