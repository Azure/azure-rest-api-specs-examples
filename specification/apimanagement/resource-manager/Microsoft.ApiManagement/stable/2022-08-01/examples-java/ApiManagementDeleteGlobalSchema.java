/** Samples for GlobalSchema Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteGlobalSchema.json
     */
    /**
     * Sample code: ApiManagementDeleteSchema.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteSchema(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .globalSchemas()
            .deleteWithResponse("rg1", "apimService1", "schema1", "*", com.azure.core.util.Context.NONE);
    }
}
