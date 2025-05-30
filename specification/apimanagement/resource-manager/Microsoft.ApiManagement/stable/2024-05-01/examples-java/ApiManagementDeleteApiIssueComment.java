
/**
 * Samples for ApiIssueComment Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteApiIssueComment.json
     */
    /**
     * Sample code: ApiManagementDeleteApiIssueComment.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteApiIssueComment(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiIssueComments().deleteWithResponse("rg1", "apimService1", "57d1f7558aa04f15146d9d8a",
            "57d2ef278aa04f0ad01d6cdc", "599e29ab193c3c0bd0b3e2fb", "*", com.azure.core.util.Context.NONE);
    }
}
