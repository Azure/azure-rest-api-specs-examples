import com.azure.core.util.Context;

/** Samples for Gateway Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetGateway.json
     */
    /**
     * Sample code: ApiManagementGetGateway.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetGateway(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.gateways().getWithResponse("rg1", "apimService1", "gw1", Context.NONE);
    }
}
