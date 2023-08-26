/** Samples for Tag GetEntityStateByProduct. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadProductTag.json
     */
    /**
     * Sample code: ApiManagementHeadProductTag.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadProductTag(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tags()
            .getEntityStateByProductWithResponse(
                "rg1",
                "apimService1",
                "59306a29e4bbd510dc24e5f8",
                "59306a29e4bbd510dc24e5f9",
                com.azure.core.util.Context.NONE);
    }
}
