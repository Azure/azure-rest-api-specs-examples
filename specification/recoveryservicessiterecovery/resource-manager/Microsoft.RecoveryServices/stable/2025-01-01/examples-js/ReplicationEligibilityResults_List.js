const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Validates whether a given VM can be protected or not in which case returns list of errors.
 *
 * @summary Validates whether a given VM can be protected or not in which case returns list of errors.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationEligibilityResults_List.json
 */
async function getsTheValidationErrorsInCaseTheVMIsUnsuitableForProtection() {
  const subscriptionId =
    process.env["RECOVERYSERVICESSITERECOVERY_SUBSCRIPTION_ID"] ||
    "d90d145a-4cdd-45a3-b2c4-971d69775278";
  const resourceGroupName = process.env["RECOVERYSERVICESSITERECOVERY_RESOURCE_GROUP"] || "testRg1";
  const virtualMachineName = "testVm2";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationEligibilityResultsOperations.list(
    resourceGroupName,
    virtualMachineName,
  );
  console.log(result);
}
