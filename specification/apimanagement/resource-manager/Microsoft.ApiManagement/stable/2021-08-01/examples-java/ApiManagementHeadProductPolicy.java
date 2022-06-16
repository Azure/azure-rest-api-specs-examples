import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/** Samples for ProductPolicy GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadProductPolicy.json
     */
    /**
     * Sample code: ApiManagementHeadProductPolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadProductPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .productPolicies()
            .getEntityTagWithResponse("rg1", "apimService1", "unlimited", PolicyIdName.POLICY, Context.NONE);
    }
}
