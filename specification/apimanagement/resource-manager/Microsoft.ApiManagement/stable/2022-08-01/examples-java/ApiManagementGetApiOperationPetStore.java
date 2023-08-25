/** Samples for ApiOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetApiOperationPetStore.json
     */
    /**
     * Sample code: ApiManagementGetApiOperationPetStore.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetApiOperationPetStore(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiOperations()
            .getWithResponse("rg1", "apimService1", "swagger-petstore", "loginUser", com.azure.core.util.Context.NONE);
    }
}
