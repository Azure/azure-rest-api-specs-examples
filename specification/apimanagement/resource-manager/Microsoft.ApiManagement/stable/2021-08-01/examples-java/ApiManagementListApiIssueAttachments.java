import com.azure.core.util.Context;

/** Samples for ApiIssueAttachment ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListApiIssueAttachments.json
     */
    /**
     * Sample code: ApiManagementListApiIssueAttachments.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListApiIssueAttachments(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiIssueAttachments()
            .listByService(
                "rg1",
                "apimService1",
                "57d1f7558aa04f15146d9d8a",
                "57d2ef278aa04f0ad01d6cdc",
                null,
                null,
                null,
                Context.NONE);
    }
}
