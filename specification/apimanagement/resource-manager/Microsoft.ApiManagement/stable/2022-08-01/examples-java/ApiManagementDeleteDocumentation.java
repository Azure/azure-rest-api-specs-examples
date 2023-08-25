/** Samples for Documentation Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteDocumentation.json
     */
    /**
     * Sample code: ApiManagementDeleteDocumentation.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteDocumentation(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .documentations()
            .deleteWithResponse(
                "rg1", "apimService1", "57d1f7558aa04f15146d9d8a", "*", com.azure.core.util.Context.NONE);
    }
}
