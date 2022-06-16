import com.azure.core.util.Context;

/** Samples for ApiManagementService Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementServiceDeleteService.json
     */
    /**
     * Sample code: ApiManagementServiceDeleteService.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementServiceDeleteService(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiManagementServices().delete("rg1", "apimService1", Context.NONE);
    }
}
