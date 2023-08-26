/** Samples for ApiPolicy ListByApi. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListApiPolicies.json
     */
    /**
     * Sample code: ApiManagementListApiPolicies.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListApiPolicies(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiPolicies()
            .listByApiWithResponse("rg1", "apimService1", "5600b59475ff190048040001", com.azure.core.util.Context.NONE);
    }
}
