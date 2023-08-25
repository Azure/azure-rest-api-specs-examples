/** Samples for ApiVersionSet GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadApiVersionSet.json
     */
    /**
     * Sample code: ApiManagementHeadApiVersionSet.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadApiVersionSet(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiVersionSets()
            .getEntityTagWithResponse("rg1", "apimService1", "vs1", com.azure.core.util.Context.NONE);
    }
}
