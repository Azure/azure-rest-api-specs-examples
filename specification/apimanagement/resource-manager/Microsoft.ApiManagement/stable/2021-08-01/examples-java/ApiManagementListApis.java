import com.azure.core.util.Context;

/** Samples for Api ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListApis.json
     */
    /**
     * Sample code: ApiManagementListApis.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListApis(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apis().listByService("rg1", "apimService1", null, null, null, null, null, Context.NONE);
    }
}
