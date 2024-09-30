
/**
 * Samples for SecurityPoliciesInterface ListByTrafficController.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/preview/2024-05-01-preview/examples/
     * SecurityPoliciesGetList.json
     */
    /**
     * Sample code: Get SecurityPolicies.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void
        getSecurityPolicies(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.securityPoliciesInterfaces().listByTrafficController("rg1", "tc1", com.azure.core.util.Context.NONE);
    }
}
