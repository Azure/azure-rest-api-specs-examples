import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/** Samples for ApiPolicy Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetApiPolicy.json
     */
    /**
     * Sample code: ApiManagementGetApiPolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetApiPolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiPolicies()
            .getWithResponse(
                "rg1",
                "apimService1",
                "5600b59475ff190048040001",
                PolicyIdName.POLICY,
                null,
                com.azure.core.util.Context.NONE);
    }
}
