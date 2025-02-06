
/**
 * Samples for SecurityPoliciesInterface ListByTrafficController.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SecurityPoliciesGetList.json
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
