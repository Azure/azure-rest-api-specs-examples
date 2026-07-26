
import com.azure.resourcemanager.hybridcompute.models.GatewayType;
import java.util.Arrays;

/**
 * Samples for Gateways CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-16-preview/gateway/Gateway_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update a Gateway.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void createOrUpdateAGateway(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.gateways().define("{gatewayName}").withRegion("eastus2euap")
            .withExistingResourceGroup("myResourceGroup").withGatewayType(GatewayType.PUBLIC)
            .withAllowedFeatures(Arrays.asList("*"))
            .withGatewayBypass(Arrays.asList("contoso.com", "internal.corp.net")).create();
    }
}
