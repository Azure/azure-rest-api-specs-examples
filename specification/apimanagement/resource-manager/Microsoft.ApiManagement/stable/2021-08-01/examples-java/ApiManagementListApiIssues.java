import com.azure.core.util.Context;

/** Samples for ApiIssue ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListApiIssues.json
     */
    /**
     * Sample code: ApiManagementListApiIssues.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListApiIssues(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiIssues()
            .listByService("rg1", "apimService1", "57d1f7558aa04f15146d9d8a", null, null, null, null, Context.NONE);
    }
}
