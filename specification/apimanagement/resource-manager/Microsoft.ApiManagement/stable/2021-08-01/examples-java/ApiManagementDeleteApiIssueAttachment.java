import com.azure.core.util.Context;

/** Samples for ApiIssueAttachment Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteApiIssueAttachment.json
     */
    /**
     * Sample code: ApiManagementDeleteApiIssueAttachment.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteApiIssueAttachment(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiIssueAttachments()
            .deleteWithResponse(
                "rg1",
                "apimService1",
                "57d1f7558aa04f15146d9d8a",
                "57d2ef278aa04f0ad01d6cdc",
                "57d2ef278aa04f0888cba3f3",
                "*",
                Context.NONE);
    }
}
