import com.azure.core.util.Context;

/** Samples for Services ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/SearchListServicesByResourceGroup.json
     */
    /**
     * Sample code: SearchListServicesByResourceGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void searchListServicesByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.searchServices().manager().serviceClient().getServices().listByResourceGroup("rg1", null, Context.NONE);
    }
}
