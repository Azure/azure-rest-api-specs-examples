
/**
 * Samples for Tag GetByApi.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetApiTag.json
     */
    /**
     * Sample code: ApiManagementGetApiTag.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetApiTag(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tags().getByApiWithResponse("rg1", "apimService1", "59d6bb8f1f7fab13dc67ec9b",
            "59306a29e4bbd510dc24e5f9", com.azure.core.util.Context.NONE);
    }
}
