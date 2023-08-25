/** Samples for Group GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadGroup.json
     */
    /**
     * Sample code: ApiManagementHeadGroup.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadGroup(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .groups()
            .getEntityTagWithResponse(
                "rg1", "apimService1", "59306a29e4bbd510dc24e5f9", com.azure.core.util.Context.NONE);
    }
}
