
import com.azure.resourcemanager.search.models.AdminKeyKind;

/**
 * Samples for AdminKeys Regenerate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/SearchRegenerateAdminKey.json
     */
    /**
     * Sample code: SearchRegenerateAdminKey.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void searchRegenerateAdminKey(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getAdminKeys().regenerateWithResponse("rg1", "mysearchservice", AdminKeyKind.PRIMARY,
            com.azure.core.util.Context.NONE);
    }
}
