/** Samples for Policies GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2022-05-01/examples/WafPolicyGet.json
     */
    /**
     * Sample code: Get Policy.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void getPolicy(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.policies().getByResourceGroupWithResponse("rg1", "Policy1", com.azure.core.util.Context.NONE);
    }
}
