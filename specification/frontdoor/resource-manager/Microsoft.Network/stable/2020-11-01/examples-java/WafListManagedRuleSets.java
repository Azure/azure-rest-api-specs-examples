/** Samples for ManagedRuleSets List. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-11-01/examples/WafListManagedRuleSets.json
     */
    /**
     * Sample code: List Policies in a Resource Group.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void listPoliciesInAResourceGroup(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.managedRuleSets().list(com.azure.core.util.Context.NONE);
    }
}
