/** Samples for ApiRelease CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiRelease.json
     */
    /**
     * Sample code: ApiManagementCreateApiRelease.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiRelease(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiReleases()
            .define("testrev")
            .withExistingApi("rg1", "apimService1", "a1")
            .withApiId(
                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/a1")
            .withNotes("yahooagain")
            .create();
    }
}
