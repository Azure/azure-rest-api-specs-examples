/** Samples for Documentation ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListDocumentations.json
     */
    /**
     * Sample code: ApiManagementListApis.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListApis(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .documentations()
            .listByService("rg1", "apimService1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
