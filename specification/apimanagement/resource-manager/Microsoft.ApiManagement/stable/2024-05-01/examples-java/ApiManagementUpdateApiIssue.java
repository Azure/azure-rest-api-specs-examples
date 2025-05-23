
import com.azure.resourcemanager.apimanagement.models.IssueContract;
import com.azure.resourcemanager.apimanagement.models.State;

/**
 * Samples for ApiIssue Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateApiIssue.json
     */
    /**
     * Sample code: ApiManagementUpdateApiIssue.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdateApiIssue(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        IssueContract resource = manager.apiIssues().getWithResponse("rg1", "apimService1", "57d1f7558aa04f15146d9d8a",
            "57d2ef278aa04f0ad01d6cdc", null, com.azure.core.util.Context.NONE).getValue();
        resource.update().withState(State.CLOSED).withIfMatch("*").apply();
    }
}
