
import com.azure.resourcemanager.containerservice.fluent.models.AgentPoolInner;
import com.azure.resourcemanager.containerservice.models.AgentPoolMode;
import com.azure.resourcemanager.containerservice.models.Code;
import com.azure.resourcemanager.containerservice.models.CreationData;
import com.azure.resourcemanager.containerservice.models.KubeletConfig;
import com.azure.resourcemanager.containerservice.models.LinuxOSConfig;
import com.azure.resourcemanager.containerservice.models.OSSku;
import com.azure.resourcemanager.containerservice.models.OSType;
import com.azure.resourcemanager.containerservice.models.PowerState;
import com.azure.resourcemanager.containerservice.models.ScaleSetEvictionPolicy;
import com.azure.resourcemanager.containerservice.models.ScaleSetPriority;
import com.azure.resourcemanager.containerservice.models.SysctlConfig;
import com.azure.resourcemanager.containerservice.models.WorkloadRuntime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AgentPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-11-01/examples/
     * AgentPoolsCreate_Update.json
     */
    /**
     * Sample code: Create/Update Agent Pool.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createUpdateAgentPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getAgentPools().createOrUpdate("rg1", "clustername1",
            "agentpool1",
            new AgentPoolInner().withCount(3).withVmSize("Standard_DS1_v2").withOsType(OSType.LINUX)
                .withMode(AgentPoolMode.USER).withOrchestratorVersion("").withScaleSetPriority(ScaleSetPriority.SPOT)
                .withScaleSetEvictionPolicy(ScaleSetEvictionPolicy.DELETE).withTags(mapOf("name1", "val1"))
                .withNodeLabels(mapOf("key1", "fakeTokenPlaceholder"))
                .withNodeTaints(Arrays.asList("Key1=Value1:NoSchedule")),
            com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-11-01/examples/
     * AgentPoolsCreate_EnableEncryptionAtHost.json
     */
    /**
     * Sample code: Create Agent Pool with EncryptionAtHost enabled.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createAgentPoolWithEncryptionAtHostEnabled(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getAgentPools().createOrUpdate(
            "rg1", "clustername1", "agentpool1", new AgentPoolInner().withCount(3).withVmSize("Standard_DS2_v2")
                .withOsType(OSType.LINUX).withOrchestratorVersion("").withEnableEncryptionAtHost(true),
            com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-11-01/examples/
     * AgentPoolsCreate_EnableUltraSSD.json
     */
    /**
     * Sample code: Create Agent Pool with UltraSSD enabled.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAgentPoolWithUltraSSDEnabled(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getAgentPools().createOrUpdate(
            "rg1", "clustername1", "agentpool1", new AgentPoolInner().withCount(3).withVmSize("Standard_DS2_v2")
                .withOsType(OSType.LINUX).withOrchestratorVersion("").withEnableUltraSsd(true),
            com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-11-01/examples/
     * AgentPoolsCreate_WasmWasi.json
     */
    /**
     * Sample code: Create Agent Pool with Krustlet and the WASI runtime.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createAgentPoolWithKrustletAndTheWASIRuntime(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getAgentPools().createOrUpdate("rg1", "clustername1",
            "agentpool1",
            new AgentPoolInner().withCount(3).withVmSize("Standard_DS2_v2").withOsDiskSizeGB(64)
                .withWorkloadRuntime(WorkloadRuntime.WASM_WASI).withOsType(OSType.LINUX).withMode(AgentPoolMode.USER)
                .withOrchestratorVersion(""),
            com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-11-01/examples/
     * AgentPoolsCreate_Snapshot.json
     */
    /**
     * Sample code: Create Agent Pool using an agent pool snapshot.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAgentPoolUsingAnAgentPoolSnapshot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getAgentPools()
            .createOrUpdate("rg1", "clustername1", "agentpool1", new AgentPoolInner().withCount(3)
                .withVmSize("Standard_DS2_v2").withOsType(OSType.LINUX).withOrchestratorVersion("").withEnableFips(true)
                .withCreationData(new CreationData().withSourceResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/snapshots/snapshot1")),
                com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-11-01/examples/
     * AgentPoolsCreate_PPG.json
     */
    /**
     * Sample code: Create Agent Pool with PPG.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAgentPoolWithPPG(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getAgentPools().createOrUpdate("rg1", "clustername1",
            "agentpool1",
            new AgentPoolInner().withCount(3).withVmSize("Standard_DS2_v2").withOsType(OSType.LINUX)
                .withOrchestratorVersion("").withProximityPlacementGroupId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.Compute/proximityPlacementGroups/ppg1"),
            com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-11-01/examples/
     * AgentPoolsCreate_CustomNodeConfig.json
     */
    /**
     * Sample code: Create Agent Pool with KubeletConfig and LinuxOSConfig.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createAgentPoolWithKubeletConfigAndLinuxOSConfig(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getAgentPools()
            .createOrUpdate("rg1", "clustername1", "agentpool1", new AgentPoolInner().withCount(3)
                .withVmSize("Standard_DS2_v2").withOsType(OSType.LINUX).withOrchestratorVersion("")
                .withKubeletConfig(new KubeletConfig().withCpuManagerPolicy("static").withCpuCfsQuota(true)
                    .withCpuCfsQuotaPeriod("200ms").withImageGcHighThreshold(90).withImageGcLowThreshold(70)
                    .withTopologyManagerPolicy("best-effort")
                    .withAllowedUnsafeSysctls(Arrays.asList("kernel.msg*", "net.core.somaxconn")).withFailSwapOn(false))
                .withLinuxOSConfig(new LinuxOSConfig()
                    .withSysctls(new SysctlConfig().withNetCoreWmemDefault(12345).withNetIpv4TcpTwReuse(true)
                        .withNetIpv4IpLocalPortRange("20000 60000").withKernelThreadsMax(99999))
                    .withTransparentHugePageEnabled("always").withTransparentHugePageDefrag("madvise")
                    .withSwapFileSizeMB(1500)),
                com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-11-01/examples/
     * AgentPools_Stop.json
     */
    /**
     * Sample code: Stop Agent Pool.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void stopAgentPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getAgentPools().createOrUpdate("rg1", "clustername1",
            "agentpool1", new AgentPoolInner().withPowerState(new PowerState().withCode(Code.STOPPED)),
            com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-11-01/examples/
     * AgentPoolsCreate_CRG.json
     */
    /**
     * Sample code: Create Agent Pool with Capacity Reservation Group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createAgentPoolWithCapacityReservationGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getAgentPools().createOrUpdate("rg1", "clustername1",
            "agentpool1",
            new AgentPoolInner().withCount(3).withVmSize("Standard_DS2_v2").withOsType(OSType.LINUX)
                .withOrchestratorVersion("").withCapacityReservationGroupId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.Compute/CapacityReservationGroups/crg1"),
            com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-11-01/examples/
     * AgentPoolsCreate_OSSKU.json
     */
    /**
     * Sample code: Create Agent Pool with OSSKU.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAgentPoolWithOSSKU(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getAgentPools().createOrUpdate("rg1", "clustername1",
            "agentpool1",
            new AgentPoolInner().withCount(3).withVmSize("Standard_DS2_v2").withOsType(OSType.LINUX)
                .withOsSku(OSSku.AZURE_LINUX).withOrchestratorVersion("")
                .withKubeletConfig(new KubeletConfig().withCpuManagerPolicy("static").withCpuCfsQuota(true)
                    .withCpuCfsQuotaPeriod("200ms").withImageGcHighThreshold(90).withImageGcLowThreshold(70)
                    .withTopologyManagerPolicy("best-effort")
                    .withAllowedUnsafeSysctls(Arrays.asList("kernel.msg*", "net.core.somaxconn")).withFailSwapOn(false))
                .withLinuxOSConfig(new LinuxOSConfig()
                    .withSysctls(new SysctlConfig().withNetCoreWmemDefault(12345).withNetIpv4TcpTwReuse(true)
                        .withNetIpv4IpLocalPortRange("20000 60000").withKernelThreadsMax(99999))
                    .withTransparentHugePageEnabled("always").withTransparentHugePageDefrag("madvise")
                    .withSwapFileSizeMB(1500)),
            com.azure.core.util.Context.NONE);
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
