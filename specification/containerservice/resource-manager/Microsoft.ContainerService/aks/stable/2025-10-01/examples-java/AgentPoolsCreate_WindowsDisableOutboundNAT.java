
import com.azure.resourcemanager.containerservice.fluent.models.AgentPoolInner;
import com.azure.resourcemanager.containerservice.models.AgentPoolWindowsProfile;
import com.azure.resourcemanager.containerservice.models.OSSku;
import com.azure.resourcemanager.containerservice.models.OSType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AgentPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * AgentPoolsCreate_WindowsDisableOutboundNAT.json
     */
    /**
     * Sample code: Create Windows Agent Pool with disabling OutboundNAT.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createWindowsAgentPoolWithDisablingOutboundNAT(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getAgentPools().createOrUpdate("rg1", "clustername1",
            "wnp2",
            new AgentPoolInner().withCount(3).withVmSize("Standard_D4s_v3").withOsType(OSType.WINDOWS)
                .withOsSku(OSSku.WINDOWS2022).withOrchestratorVersion("1.23.8")
                .withWindowsProfile(new AgentPoolWindowsProfile().withDisableOutboundNat(true)),
            null, null, com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
