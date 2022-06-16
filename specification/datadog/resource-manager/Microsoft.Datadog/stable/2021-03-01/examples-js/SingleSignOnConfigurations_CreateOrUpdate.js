const { MicrosoftDatadogClient } = require("@azure/arm-datadog");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Configures single-sign-on for this resource.
 *
 * @summary Configures single-sign-on for this resource.
 * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/SingleSignOnConfigurations_CreateOrUpdate.json
 */
async function singleSignOnConfigurationsCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const configurationName = "default";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftDatadogClient(credential, subscriptionId);
  const result = await client.singleSignOnConfigurations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    monitorName,
    configurationName
  );
  console.log(result);
}

singleSignOnConfigurationsCreateOrUpdate().catch(console.error);
