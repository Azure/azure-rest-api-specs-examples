
/**
 * Samples for SecurityPoliciesInterface Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/preview/2024-05-01-preview/examples/
     * SecurityPolicyGet.json
     */
    /**
     * Sample code: Get SecurityPolicy.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void getSecurityPolicy(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.securityPoliciesInterfaces().getWithResponse("rg1", "tc1", "sp1", com.azure.core.util.Context.NONE);
    }
}
