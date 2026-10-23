
/**
 * Samples for AgentFabricMemberships Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-02-preview/AgentFabricMemberships_Get.json
     */
    /**
     * Sample code: Get AgentFabric Membership.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        getAgentFabricMembership(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentFabricMemberships().getWithResponse("rg1", "clustername1", "default",
            com.azure.core.util.Context.NONE);
    }
}
