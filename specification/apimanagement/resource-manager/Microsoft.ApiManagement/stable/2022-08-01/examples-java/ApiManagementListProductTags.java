/** Samples for Tag ListByProduct. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListProductTags.json
     */
    /**
     * Sample code: ApiManagementListProductTags.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListProductTags(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tags()
            .listByProduct(
                "rg1", "apimService1", "57d2ef278aa04f0888cba3f1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
