
/**
 * Samples for ApiIssueAttachment CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateApiIssueAttachment.json
     */
    /**
     * Sample code: ApiManagementCreateApiIssueAttachment.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateApiIssueAttachment(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiIssueAttachments().define("57d2ef278aa04f0888cba3f3")
            .withExistingIssue("rg1", "apimService1", "57d1f7558aa04f15146d9d8a", "57d2ef278aa04f0ad01d6cdc")
            .withTitle("Issue attachment.").withContentFormat("image/jpeg").withContent("IEJhc2U2NA==").create();
    }
}
