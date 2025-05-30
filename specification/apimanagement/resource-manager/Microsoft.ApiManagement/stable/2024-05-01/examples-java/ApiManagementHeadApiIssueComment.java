
/**
 * Samples for ApiIssueComment GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadApiIssueComment.json
     */
    /**
     * Sample code: ApiManagementHeadApiIssueComment.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadApiIssueComment(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiIssueComments().getEntityTagWithResponse("rg1", "apimService1", "57d2ef278aa04f0888cba3f3",
            "57d2ef278aa04f0ad01d6cdc", "599e29ab193c3c0bd0b3e2fb", com.azure.core.util.Context.NONE);
    }
}
