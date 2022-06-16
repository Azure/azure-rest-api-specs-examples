/** Samples for Tag CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateTag.json
     */
    /**
     * Sample code: ApiManagementCreateTag.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateTag(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tags().define("tagId1").withExistingService("rg1", "apimService1").withDisplayName("tag1").create();
    }
}
