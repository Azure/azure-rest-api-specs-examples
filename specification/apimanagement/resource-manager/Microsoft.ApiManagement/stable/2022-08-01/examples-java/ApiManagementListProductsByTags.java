/** Samples for Product ListByTags. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListProductsByTags.json
     */
    /**
     * Sample code: ApiManagementListProductsByTags.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListProductsByTags(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.products().listByTags("rg1", "apimService1", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
