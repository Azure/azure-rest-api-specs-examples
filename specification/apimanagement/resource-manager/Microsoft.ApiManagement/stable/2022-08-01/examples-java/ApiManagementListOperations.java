/** Samples for ApiManagementOperations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListOperations.json
     */
    /**
     * Sample code: ApiManagementListOperations.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListOperations(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiManagementOperations().list(com.azure.core.util.Context.NONE);
    }
}
