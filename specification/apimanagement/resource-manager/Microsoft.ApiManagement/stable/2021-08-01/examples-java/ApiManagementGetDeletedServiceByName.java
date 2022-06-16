import com.azure.core.util.Context;

/** Samples for DeletedServices GetByName. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetDeletedServiceByName.json
     */
    /**
     * Sample code: ApiManagementGetDeletedServiceByName.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetDeletedServiceByName(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.deletedServices().getByNameWithResponse("apimService3", "westus", Context.NONE);
    }
}
