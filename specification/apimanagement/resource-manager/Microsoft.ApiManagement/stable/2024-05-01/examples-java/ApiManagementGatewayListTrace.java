
import com.azure.resourcemanager.apimanagement.models.GatewayListTraceContract;

/**
 * Samples for Gateway ListTrace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGatewayListTrace.json
     */
    /**
     * Sample code: ApiManagementGatewayListTrace.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGatewayListTrace(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.gateways().listTraceWithResponse("rg1", "apimService1", "gw1",
            new GatewayListTraceContract().withTraceId("CrDvXXXXXXXXXXXXXVU3ZA2-1"), com.azure.core.util.Context.NONE);
    }
}
