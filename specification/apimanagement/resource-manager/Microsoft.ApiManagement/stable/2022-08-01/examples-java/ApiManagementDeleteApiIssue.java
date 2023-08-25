/** Samples for ApiIssue Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteApiIssue.json
     */
    /**
     * Sample code: ApiManagementDeleteApiIssue.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteApiIssue(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiIssues()
            .deleteWithResponse(
                "rg1",
                "apimService1",
                "57d1f7558aa04f15146d9d8a",
                "57d2ef278aa04f0ad01d6cdc",
                "*",
                com.azure.core.util.Context.NONE);
    }
}
