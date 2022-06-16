import com.azure.core.util.Context;

/** Samples for ApiIssueAttachment Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiIssueAttachment.json
     */
    /**
     * Sample code: ApiManagementGetApiIssueAttachment.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetApiIssueAttachment(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiIssueAttachments()
            .getWithResponse(
                "rg1",
                "apimService1",
                "57d2ef278aa04f0888cba3f3",
                "57d2ef278aa04f0ad01d6cdc",
                "57d2ef278aa04f0888cba3f3",
                Context.NONE);
    }
}
