import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/** Samples for Policy Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetPolicy.json
     */
    /**
     * Sample code: ApiManagementGetPolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetPolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .policies()
            .getWithResponse("rg1", "apimService1", PolicyIdName.POLICY, null, com.azure.core.util.Context.NONE);
    }
}
