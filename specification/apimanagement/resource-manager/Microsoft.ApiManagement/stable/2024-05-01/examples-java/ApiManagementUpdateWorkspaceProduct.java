
import com.azure.resourcemanager.apimanagement.models.ProductUpdateParameters;

/**
 * Samples for WorkspaceProduct Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateWorkspaceProduct.json
     */
    /**
     * Sample code: ApiManagementUpdateWorkspaceProduct.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdateWorkspaceProduct(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceProducts().updateWithResponse("rg1", "apimService1", "wks1", "testproduct", "*",
            new ProductUpdateParameters().withDisplayName("Test Template ProductName 4"),
            com.azure.core.util.Context.NONE);
    }
}
