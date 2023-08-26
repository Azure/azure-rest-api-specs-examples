/** Samples for GlobalSchema Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetGlobalSchema2.json
     */
    /**
     * Sample code: ApiManagementGetSchema2.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetSchema2(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.globalSchemas().getWithResponse("rg1", "apimService1", "schema2", com.azure.core.util.Context.NONE);
    }
}
