import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.GatewayContract;
import com.azure.resourcemanager.apimanagement.models.ResourceLocationDataContract;

/** Samples for Gateway Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateGateway.json
     */
    /**
     * Sample code: ApiManagementUpdateGateway.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateGateway(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        GatewayContract resource =
            manager.gateways().getWithResponse("rg1", "apimService1", "gw1", Context.NONE).getValue();
        resource
            .update()
            .withLocationData(new ResourceLocationDataContract().withName("my location"))
            .withDescription("my gateway 1")
            .withIfMatch("*")
            .apply();
    }
}
