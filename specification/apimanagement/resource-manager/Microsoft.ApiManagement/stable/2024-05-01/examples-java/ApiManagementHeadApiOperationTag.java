
/**
 * Samples for Tag GetEntityStateByOperation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadApiOperationTag.json
     */
    /**
     * Sample code: ApiManagementHeadApiOperationTag.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadApiOperationTag(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tags().getEntityStateByOperationWithResponse("rg1", "apimService1", "59d6bb8f1f7fab13dc67ec9b",
            "59d6bb8f1f7fab13dc67ec9a", "59306a29e4bbd510dc24e5f9", com.azure.core.util.Context.NONE);
    }
}
