import com.azure.core.util.Context;

/** Samples for Tag GetEntityStateByApi. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadApiTag.json
     */
    /**
     * Sample code: ApiManagementHeadApiTag.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadApiTag(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tags()
            .getEntityStateByApiWithResponse(
                "rg1", "apimService1", "59d6bb8f1f7fab13dc67ec9b", "59306a29e4bbd510dc24e5f9", Context.NONE);
    }
}
