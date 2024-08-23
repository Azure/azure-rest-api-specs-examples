
import com.azure.resourcemanager.search.models.AdminKeyKind;

/**
 * Samples for AdminKeys Regenerate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/SearchRegenerateAdminKey.json
     */
    /**
     * Sample code: SearchRegenerateAdminKey.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void searchRegenerateAdminKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.searchServices().manager().serviceClient().getAdminKeys().regenerateWithResponse("rg1", "mysearchservice",
            AdminKeyKind.PRIMARY, null, com.azure.core.util.Context.NONE);
    }
}
