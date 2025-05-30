
import com.azure.resourcemanager.apimanagement.models.ApiReleaseContract;

/**
 * Samples for ApiRelease Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateApiRelease.json
     */
    /**
     * Sample code: ApiManagementUpdateApiRelease.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdateApiRelease(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        ApiReleaseContract resource = manager.apiReleases()
            .getWithResponse("rg1", "apimService1", "a1", "testrev", com.azure.core.util.Context.NONE).getValue();
        resource.update().withApiId(
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/a1")
            .withNotes("yahooagain").withIfMatch("*").apply();
    }
}
