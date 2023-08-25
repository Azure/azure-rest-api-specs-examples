/** Samples for GlobalSchema Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetGlobalSchema1.json
     */
    /**
     * Sample code: ApiManagementGetSchema1.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetSchema1(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.globalSchemas().getWithResponse("rg1", "apimService1", "schema1", com.azure.core.util.Context.NONE);
    }
}
