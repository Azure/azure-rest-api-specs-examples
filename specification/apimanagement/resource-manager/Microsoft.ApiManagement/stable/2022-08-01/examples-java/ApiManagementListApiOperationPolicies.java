/** Samples for ApiOperationPolicy ListByOperation. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListApiOperationPolicies.json
     */
    /**
     * Sample code: ApiManagementListApiOperationPolicies.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListApiOperationPolicies(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiOperationPolicies()
            .listByOperationWithResponse(
                "rg1",
                "apimService1",
                "599e2953193c3c0bd0b3e2fa",
                "599e29ab193c3c0bd0b3e2fb",
                com.azure.core.util.Context.NONE);
    }
}
