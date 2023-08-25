/** Samples for Group CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateGroup.json
     */
    /**
     * Sample code: ApiManagementCreateGroup.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateGroup(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .groups()
            .define("tempgroup")
            .withExistingService("rg1", "apimService1")
            .withDisplayName("temp group")
            .create();
    }
}
