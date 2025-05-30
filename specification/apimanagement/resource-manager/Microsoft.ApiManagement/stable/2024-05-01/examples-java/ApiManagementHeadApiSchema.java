
/**
 * Samples for ApiSchema GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadApiSchema.json
     */
    /**
     * Sample code: ApiManagementHeadApiSchema.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadApiSchema(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiSchemas().getEntityTagWithResponse("rg1", "apimService1", "57d1f7558aa04f15146d9d8a",
            "ec12520d-9d48-4e7b-8f39-698ca2ac63f1", com.azure.core.util.Context.NONE);
    }
}
