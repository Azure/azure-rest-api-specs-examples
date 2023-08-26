/** Samples for ApiIssue Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetApiIssue.json
     */
    /**
     * Sample code: ApiManagementGetApiIssue.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetApiIssue(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiIssues()
            .getWithResponse(
                "rg1",
                "apimService1",
                "57d2ef278aa04f0888cba3f3",
                "57d2ef278aa04f0ad01d6cdc",
                null,
                com.azure.core.util.Context.NONE);
    }
}
