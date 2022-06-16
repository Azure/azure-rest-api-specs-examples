const { MicrosoftDatadogClient } = require("@azure/arm-datadog");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the single sign-on configurations for a given monitor resource.
 *
 * @summary List the single sign-on configurations for a given monitor resource.
 * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/SingleSignOnConfigurations_List.json
 */
async function singleSignOnConfigurationsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftDatadogClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.singleSignOnConfigurations.list(resourceGroupName, monitorName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

singleSignOnConfigurationsList().catch(console.error);
