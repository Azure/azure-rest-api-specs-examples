import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/** Samples for Policy GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadPolicy.json
     */
    /**
     * Sample code: ApiManagementHeadPolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadPolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .policies()
            .getEntityTagWithResponse("rg1", "apimService1", PolicyIdName.POLICY, com.azure.core.util.Context.NONE);
    }
}
