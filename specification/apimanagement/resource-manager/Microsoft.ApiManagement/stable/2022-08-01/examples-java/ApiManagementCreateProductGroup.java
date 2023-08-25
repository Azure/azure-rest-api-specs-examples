/** Samples for ProductGroup CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateProductGroup.json
     */
    /**
     * Sample code: ApiManagementCreateProductGroup.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateProductGroup(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .productGroups()
            .createOrUpdateWithResponse(
                "rg1", "apimService1", "testproduct", "templateGroup", com.azure.core.util.Context.NONE);
    }
}
