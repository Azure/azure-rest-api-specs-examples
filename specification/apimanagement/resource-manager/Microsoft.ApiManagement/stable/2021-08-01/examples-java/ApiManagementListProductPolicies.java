import com.azure.core.util.Context;

/** Samples for ProductPolicy ListByProduct. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListProductPolicies.json
     */
    /**
     * Sample code: ApiManagementListProductPolicies.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListProductPolicies(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.productPolicies().listByProductWithResponse("rg1", "apimService1", "armTemplateProduct4", Context.NONE);
    }
}
