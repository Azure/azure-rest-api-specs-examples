/** Samples for Policies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2022-05-01/examples/WafListPoliciesUnderSubscription.json
     */
    /**
     * Sample code: Get all Policies in a Resource Group.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void getAllPoliciesInAResourceGroup(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.policies().list(com.azure.core.util.Context.NONE);
    }
}
