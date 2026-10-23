
/**
 * Samples for AgentFabricMemberships Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-02-preview/AgentFabricMemberships_Delete.json
     */
    /**
     * Sample code: Delete AgentFabric Membership.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        deleteAgentFabricMembership(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentFabricMemberships().delete("rg1", "clustername1", "default",
            com.azure.core.util.Context.NONE);
    }
}
