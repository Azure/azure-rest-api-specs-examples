
import com.azure.resourcemanager.containerservice.fluent.models.AgentFabricMembershipInner;
import com.azure.resourcemanager.containerservice.models.AgentFabricMembershipProperties;
import com.azure.resourcemanager.containerservice.models.AgentFabricMembershipSkuName;

/**
 * Samples for AgentFabricMemberships CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-02-preview/AgentFabricMemberships_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update AgentFabric Membership.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void createOrUpdateAgentFabricMembership(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentFabricMemberships().createOrUpdate("rg1", "clustername1", "default",
            new AgentFabricMembershipInner()
                .withProperties(new AgentFabricMembershipProperties().withSku(AgentFabricMembershipSkuName.STANDARD))
                .withManagedBy(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.NetworkSecurity/agentFabrics/fabric1"),
            com.azure.core.util.Context.NONE);
    }
}
