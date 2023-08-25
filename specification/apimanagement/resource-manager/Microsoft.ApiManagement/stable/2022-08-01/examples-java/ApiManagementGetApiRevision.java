/** Samples for Api Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetApiRevision.json
     */
    /**
     * Sample code: ApiManagementGetApiRevisionContract.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetApiRevisionContract(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apis().getWithResponse("rg1", "apimService1", "echo-api;rev=3", com.azure.core.util.Context.NONE);
    }
}
