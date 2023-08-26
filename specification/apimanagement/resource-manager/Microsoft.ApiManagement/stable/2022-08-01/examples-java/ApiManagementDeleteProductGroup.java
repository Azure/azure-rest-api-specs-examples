/** Samples for ProductGroup Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteProductGroup.json
     */
    /**
     * Sample code: ApiManagementDeleteProductGroup.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteProductGroup(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .productGroups()
            .deleteWithResponse(
                "rg1", "apimService1", "testproduct", "templateGroup", com.azure.core.util.Context.NONE);
    }
}
