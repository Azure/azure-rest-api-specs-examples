
import com.azure.resourcemanager.appservice.fluent.models.SiteInner;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for WebApps CreateOrUpdateSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/CreateOrUpdateWebAppSlot.json
     */
    /**
     * Sample code: Create or Update Web App Slot.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void createOrUpdateWebAppSlot(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().createOrUpdateSlot("testrg123", "sitef6141", "staging",
            new SiteInner().withLocation("East US").withKind("app").withServerFarmId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/serverfarms/DefaultAsp"),
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
