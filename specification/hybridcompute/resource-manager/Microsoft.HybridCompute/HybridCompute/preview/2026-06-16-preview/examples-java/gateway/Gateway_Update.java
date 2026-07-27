
import com.azure.resourcemanager.hybridcompute.models.Gateway;
import java.util.Arrays;

/**
 * Samples for Gateways Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-16-preview/gateway/Gateway_Update.json
     */
    /**
     * Sample code: Update a Gateway.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void updateAGateway(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        Gateway resource = manager.gateways()
            .getByResourceGroupWithResponse("myResourceGroup", "{gatewayName}", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withAllowedFeatures(Arrays.asList("*"))
            .withGatewayBypass(Arrays.asList("contoso.com", "internal.corp.net")).apply();
    }
}
