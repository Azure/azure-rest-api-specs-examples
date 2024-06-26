const { NginxManagementClient } = require("@azure/arm-nginx");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Analyze an NGINX configuration without applying it to the NGINXaaS deployment
 *
 * @summary Analyze an NGINX configuration without applying it to the NGINXaaS deployment
 * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-01-01-preview/examples/Configurations_Analysis.json
 */
async function configurationsAnalysis() {
  const subscriptionId =
    process.env["NGINX_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["NGINX_RESOURCE_GROUP"] || "myResourceGroup";
  const deploymentName = "myDeployment";
  const configurationName = "default";
  const body = {
    config: {
      files: [{ content: "ABCDEF==", virtualPath: "/etc/nginx/nginx.conf" }],
      package: { data: undefined },
      rootFile: "/etc/nginx/nginx.conf",
    },
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new NginxManagementClient(credential, subscriptionId);
  const result = await client.configurations.analysis(
    resourceGroupName,
    deploymentName,
    configurationName,
    options,
  );
  console.log(result);
}
