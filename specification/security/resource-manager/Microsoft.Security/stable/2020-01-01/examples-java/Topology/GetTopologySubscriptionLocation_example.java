import com.azure.core.util.Context;

/** Samples for Topology ListByHomeRegion. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/Topology/GetTopologySubscriptionLocation_example.json
     */
    /**
     * Sample code: Get topology on a subscription from security data location.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getTopologyOnASubscriptionFromSecurityDataLocation(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.topologies().listByHomeRegion("centralus", Context.NONE);
    }
}
