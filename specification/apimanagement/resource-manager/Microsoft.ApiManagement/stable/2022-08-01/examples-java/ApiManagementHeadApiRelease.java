/** Samples for ApiRelease GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadApiRelease.json
     */
    /**
     * Sample code: ApiManagementHeadApiRelease.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadApiRelease(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiReleases()
            .getEntityTagWithResponse(
                "rg1", "apimService1", "a1", "5a7cb545298324c53224a799", com.azure.core.util.Context.NONE);
    }
}
