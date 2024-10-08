const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates applications for the HDInsight cluster.
 *
 * @summary Creates applications for the HDInsight cluster.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/CreateApplication.json
 */
async function createApplication() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cluster1";
  const applicationName = "hue";
  const parameters = {
    properties: {
      applicationType: "CustomApplication",
      computeProfile: {
        roles: [
          {
            name: "edgenode",
            hardwareProfile: { vmSize: "Standard_D12_v2" },
            targetInstanceCount: 1,
          },
        ],
      },
      errors: [],
      httpsEndpoints: [
        {
          accessModes: ["WebPage"],
          destinationPort: 20000,
          subDomainSuffix: "dss",
        },
      ],
      installScriptActions: [
        {
          name: "app-install-app1",
          parameters: "-version latest -port 20000",
          roles: ["edgenode"],
          uri: "https://.../install.sh",
        },
      ],
      uninstallScriptActions: [],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.applications.beginCreateAndWait(
    resourceGroupName,
    clusterName,
    applicationName,
    parameters,
  );
  console.log(result);
}
