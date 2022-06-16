import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.PolicyScopeContract;

/** Samples for PolicyDescription ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListPolicyDescriptions.json
     */
    /**
     * Sample code: ApiManagementListPolicyDescriptions.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListPolicyDescriptions(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .policyDescriptions()
            .listByServiceWithResponse("rg1", "apimService1", PolicyScopeContract.API, Context.NONE);
    }
}
