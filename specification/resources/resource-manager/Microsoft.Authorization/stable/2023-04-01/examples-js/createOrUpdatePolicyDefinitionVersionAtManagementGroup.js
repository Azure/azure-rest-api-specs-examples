const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation creates or updates a policy definition version in the given management group with the given name.
 *
 * @summary This operation creates or updates a policy definition version in the given management group with the given name.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/createOrUpdatePolicyDefinitionVersionAtManagementGroup.json
 */
async function createOrUpdateAPolicyDefinitionVersionAtManagementGroupLevel() {
  const managementGroupName = "MyManagementGroup";
  const policyDefinitionName = "ResourceNaming";
  const policyDefinitionVersion = "1.2.1";
  const parameters = {
    description: "Force resource names to begin with given 'prefix' and/or end with given 'suffix'",
    displayName: "Enforce resource naming convention",
    metadata: { category: "Naming" },
    mode: "All",
    parameters: {
      prefix: {
        type: "String",
        metadata: {
          description: "Resource name prefix",
          displayName: "Prefix",
        },
      },
      suffix: {
        type: "String",
        metadata: {
          description: "Resource name suffix",
          displayName: "Suffix",
        },
      },
    },
    policyRule: {
      if: {
        not: {
          field: "name",
          like: "[concat(parameters('prefix'), '*', parameters('suffix'))]",
        },
      },
      then: { effect: "deny" },
    },
    version: "1.2.1",
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.policyDefinitionVersions.createOrUpdateAtManagementGroup(
    managementGroupName,
    policyDefinitionName,
    policyDefinitionVersion,
    parameters,
  );
  console.log(result);
}
