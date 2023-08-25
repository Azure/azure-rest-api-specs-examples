import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/** Samples for Policy Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeletePolicy.json
     */
    /**
     * Sample code: ApiManagementDeletePolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeletePolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .policies()
            .deleteWithResponse("rg1", "apimService1", PolicyIdName.POLICY, "*", com.azure.core.util.Context.NONE);
    }
}
