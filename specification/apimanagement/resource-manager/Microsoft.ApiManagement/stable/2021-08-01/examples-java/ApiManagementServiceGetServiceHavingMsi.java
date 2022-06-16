import com.azure.core.util.Context;

/** Samples for ApiManagementService GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementServiceGetServiceHavingMsi.json
     */
    /**
     * Sample code: ApiManagementServiceGetServiceHavingMsi.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementServiceGetServiceHavingMsi(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiManagementServices().getByResourceGroupWithResponse("rg1", "apimService1", Context.NONE);
    }
}
