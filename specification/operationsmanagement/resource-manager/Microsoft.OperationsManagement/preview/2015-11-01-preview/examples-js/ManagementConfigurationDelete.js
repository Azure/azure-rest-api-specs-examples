const { OperationsManagementClient } = require("@azure/arm-operations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the ManagementConfiguration in the subscription.
 *
 * @summary Deletes the ManagementConfiguration in the subscription.
 * x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/ManagementConfigurationDelete.json
 */
async function managementConfigurationDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const managementConfigurationName = "managementConfigurationName";
  const credential = new DefaultAzureCredential();
  const client = new OperationsManagementClient(credential, subscriptionId);
  const result = await client.managementConfigurations.delete(
    resourceGroupName,
    managementConfigurationName
  );
  console.log(result);
}

managementConfigurationDelete().catch(console.error);
