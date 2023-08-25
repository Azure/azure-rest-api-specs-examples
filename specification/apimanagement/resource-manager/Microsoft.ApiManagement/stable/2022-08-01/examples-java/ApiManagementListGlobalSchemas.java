/** Samples for GlobalSchema ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListGlobalSchemas.json
     */
    /**
     * Sample code: ApiManagementListSchemas.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListSchemas(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .globalSchemas()
            .listByService("rg1", "apimService1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
