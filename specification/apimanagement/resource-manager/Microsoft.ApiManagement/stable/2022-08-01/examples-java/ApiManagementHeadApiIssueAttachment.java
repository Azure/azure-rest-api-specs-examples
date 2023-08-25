/** Samples for ApiIssueAttachment GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadApiIssueAttachment.json
     */
    /**
     * Sample code: ApiManagementHeadApiIssueAttachment.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadApiIssueAttachment(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiIssueAttachments()
            .getEntityTagWithResponse(
                "rg1",
                "apimService1",
                "57d2ef278aa04f0888cba3f3",
                "57d2ef278aa04f0ad01d6cdc",
                "57d2ef278aa04f0888cba3f3",
                com.azure.core.util.Context.NONE);
    }
}
