
/**
 * Samples for Topology Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/Topology/
     * GetTopology_example.json
     */
    /**
     * Sample code: Get topology.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getTopology(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.topologies().getWithResponse("myservers", "centralus", "vnets", com.azure.core.util.Context.NONE);
    }
}
