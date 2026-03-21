const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an agent pool in the specified managed cluster.
 *
 * @summary creates or updates an agent pool in the specified managed cluster.
 * x-ms-original-file: 2026-01-01/AgentPoolsCreate_GPUMIG.json
 */
async function createAgentPoolWithGpumig() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.agentPools.createOrUpdate("rg1", "clustername1", "agentpool1", {
    count: 3,
    gpuInstanceProfile: "MIG2g",
    kubeletConfig: {
      allowedUnsafeSysctls: ["kernel.msg*", "net.core.somaxconn"],
      cpuCfsQuota: true,
      cpuCfsQuotaPeriod: "200ms",
      cpuManagerPolicy: "static",
      failSwapOn: false,
      imageGcHighThreshold: 90,
      imageGcLowThreshold: 70,
      topologyManagerPolicy: "best-effort",
    },
    linuxOSConfig: {
      swapFileSizeMB: 1500,
      sysctls: {
        kernelThreadsMax: 99999,
        netCoreWmemDefault: 12345,
        netIpv4IpLocalPortRange: "20000 60000",
        netIpv4TcpTwReuse: true,
      },
      transparentHugePageDefrag: "madvise",
      transparentHugePageEnabled: "always",
    },
    orchestratorVersion: "",
    osType: "Linux",
    vmSize: "Standard_ND96asr_v4",
  });
  console.log(result);
}
