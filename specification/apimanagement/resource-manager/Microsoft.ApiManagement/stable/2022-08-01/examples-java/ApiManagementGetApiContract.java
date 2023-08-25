/** Samples for Api Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetApiContract.json
     */
    /**
     * Sample code: ApiManagementGetApiContract.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetApiContract(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apis()
            .getWithResponse("rg1", "apimService1", "57d1f7558aa04f15146d9d8a", com.azure.core.util.Context.NONE);
    }
}
