Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.containerservice.fluent.models.AgentPoolInner;
import com.azure.resourcemanager.containerservice.models.GpuInstanceProfile;
import com.azure.resourcemanager.containerservice.models.KubeletConfig;
import com.azure.resourcemanager.containerservice.models.LinuxOSConfig;
import com.azure.resourcemanager.containerservice.models.OSType;
import com.azure.resourcemanager.containerservice.models.SysctlConfig;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for AgentPools CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/AgentPoolsCreate_GPUMIG.json
     */
    /**
     * Sample code: Create Agent Pool with GPUMIG.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAgentPoolWithGPUMIG(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getAgentPools()
            .createOrUpdate(
                "rg1",
                "clustername1",
                "agentpool1",
                new AgentPoolInner()
                    .withCount(3)
                    .withVmSize("Standard_ND96asr_v4")
                    .withOsType(OSType.LINUX)
                    .withOrchestratorVersion("")
                    .withKubeletConfig(
                        new KubeletConfig()
                            .withCpuManagerPolicy("static")
                            .withCpuCfsQuota(true)
                            .withCpuCfsQuotaPeriod("200ms")
                            .withImageGcHighThreshold(90)
                            .withImageGcLowThreshold(70)
                            .withTopologyManagerPolicy("best-effort")
                            .withAllowedUnsafeSysctls(Arrays.asList("kernel.msg*", "net.core.somaxconn"))
                            .withFailSwapOn(false))
                    .withLinuxOSConfig(
                        new LinuxOSConfig()
                            .withSysctls(
                                new SysctlConfig()
                                    .withNetCoreWmemDefault(12345)
                                    .withNetIpv4TcpTwReuse(true)
                                    .withNetIpv4IpLocalPortRange("20000 60000")
                                    .withKernelThreadsMax(99999))
                            .withTransparentHugePageEnabled("always")
                            .withTransparentHugePageDefrag("madvise")
                            .withSwapFileSizeMB(1500))
                    .withGpuInstanceProfile(GpuInstanceProfile.MIG2G),
                Context.NONE);
    }

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
```
