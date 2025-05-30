
/**
 * Samples for ProductGroupLink CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateProductGroupLink.json
     */
    /**
     * Sample code: ApiManagementCreateProductGroupLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateProductGroupLink(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.productGroupLinks().define("link1").withExistingProduct("rg1", "apimService1", "testproduct")
            .withGroupId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/groups/group1")
            .create();
    }
}
