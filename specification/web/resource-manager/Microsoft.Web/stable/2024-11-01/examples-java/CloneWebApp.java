
import com.azure.resourcemanager.appservice.fluent.models.SiteInner;
import com.azure.resourcemanager.appservice.models.CloningInfo;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for WebApps CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/CloneWebApp.json
     */
    /**
     * Sample code: Clone web app.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cloneWebApp(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().createOrUpdate("testrg123", "sitef6141",
            new SiteInner().withLocation("East US").withKind("app").withCloningInfo(new CloningInfo()
                .withOverwrite(false).withCloneCustomHostNames(true).withCloneSourceControl(true)
                .withSourceWebAppId(
                    "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg456/providers/Microsoft.Web/sites/srcsiteg478")
                .withSourceWebAppLocation("West Europe")
                .withHostingEnvironment(
                    "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg456/providers/Microsoft.Web/hostingenvironments/aseforsites")
                .withAppSettingsOverrides(mapOf("Setting1", "NewValue1", "Setting3", "NewValue5"))
                .withConfigureLoadBalancing(false)),
            com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
