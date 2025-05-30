
/**
 * Samples for ApiSchema Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteApiSchema.json
     */
    /**
     * Sample code: ApiManagementDeleteApiSchema.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteApiSchema(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiSchemas().deleteWithResponse("rg1", "apimService1", "59d5b28d1f7fab116c282650",
            "59d5b28e1f7fab116402044e", "*", null, com.azure.core.util.Context.NONE);
    }
}
