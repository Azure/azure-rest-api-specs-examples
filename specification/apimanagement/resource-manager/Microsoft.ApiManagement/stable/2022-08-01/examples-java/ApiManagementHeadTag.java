/** Samples for Tag GetEntityState. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadTag.json
     */
    /**
     * Sample code: ApiManagementHeadTag.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadTag(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tags()
            .getEntityStateWithResponse(
                "rg1", "apimService1", "59306a29e4bbd510dc24e5f9", com.azure.core.util.Context.NONE);
    }
}
