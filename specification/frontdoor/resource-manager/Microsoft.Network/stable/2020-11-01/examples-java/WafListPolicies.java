/** Samples for Policies ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-11-01/examples/WafListPolicies.json
     */
    /**
     * Sample code: List Policies in a Resource Group.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void listPoliciesInAResourceGroup(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.policies().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
