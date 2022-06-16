import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/** Samples for ProductPolicy Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetProductPolicy.json
     */
    /**
     * Sample code: ApiManagementGetProductPolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetProductPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .productPolicies()
            .getWithResponse(
                "rg1", "apimService1", "kjoshiarmTemplateProduct4", PolicyIdName.POLICY, null, Context.NONE);
    }
}
