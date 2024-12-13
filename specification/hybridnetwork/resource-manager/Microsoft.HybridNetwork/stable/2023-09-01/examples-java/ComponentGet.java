
/**
 * Samples for Components Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ComponentGet.json
     */
    /**
     * Sample code: Get component resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getComponentResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.components().getWithResponse("rg", "testNf", "testComponent", com.azure.core.util.Context.NONE);
    }
}
